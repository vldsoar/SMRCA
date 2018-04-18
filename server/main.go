package main

import (
	"net"
	"fmt"
	"os"
	"encoding/json"
	"errors"
	"log"
	"io"
	"bytes"
	"reflect"
	"client-server/server/db"
	"flag"
	"time"
	"client-server/server/mail"
)

type Request struct {
	Method string `json:"method"`
	Action string `json:"action"`
	Params map[string]interface{} `json:"params,omitempty"`
}

type Response struct {
	Success bool                   `json:"success"`
	Error   string                 `json:"error,omitempty"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

type Model interface {
	save()
}

type ConsumptionInfo struct {
	SensorID int `json:"sensor_id"`
	Date string `json:"date"`
	Measure uint64 `json:"measure"`
}

//func (c *ConsumptionInfo) save2() bool {
//	j, err := db.Load("sensors")
//
//	if err != nil {
//		return false
//	}
//
//	jSensor, err := db.GetSensor(c.SensorID)
//
//	if err != nil {
//		fmt.Println(err)
//		return false
//	}
//
//	measures, _ := jSensor.Get("measures").Array()
//
//	updatedMeasures := append(measures, c)
//
//	jSensor.Set("measures", updatedMeasures)
//
//	fmt.Println(jSensor.MustString())
//
//	dir, _ := os.Getwd()
//
//	f, err := os.Create(path.Join(dir, "server/db/sensors.json"))
//
//	if err != nil {
//		fmt.Println(err)
//		return false
//	}
//
//	jEncode, _ := json.MarshalIndent(j,"", "\t")
//
//	_, err = io.Copy(f, bytes.NewReader(jEncode))
//
//	if err != nil {
//		fmt.Println(err)
//		return false
//	}
//
//	return true
//	//fmt.Println(reflect.TypeOf(x).Kind())
//}

func (c *ConsumptionInfo) save() bool {
	j, err := db.Load("sensors")

	if err != nil {
		return false
	}

	data, err := j.Array()

	if err != nil {
		fmt.Println(err)
		return false
	}

	save := false

	for i := range data {
		wantedSensor := data[i].(map[string]interface{})

		sensorID, _ := wantedSensor["id"].(json.Number).Int64()

		if sensorID == int64(c.SensorID) {
			sensorFound := j.GetIndex(i)

			measures, _ := sensorFound.Get("measures").Array()

			updatedMeasures := append(measures, c)

			sensorFound.Set("measures", updatedMeasures)

			f, err := os.Create("server/db/sensors.json")

			if err != nil {
				fmt.Println(err)
				break
			}

			jEncode, _ := json.MarshalIndent(j,"", "\t")

			_, err = io.Copy(f, bytes.NewReader(jEncode))

			if err != nil {
				fmt.Println(err)
				break
			}

			save = true
		}
	}

	//fmt.Println(reflect.TypeOf(x).Kind())

	return save

}

func main() {

	ip, _ := externalIP()

	localIP := flag.String("ip", ip, "192.0.0.1")
	portTCP := flag.String("portTCP", "8000", "8000")
	portUDP := flag.String("portUDP", "8001", "8001")

	flag.Parse()

	// Start the TCP server
	go startServerTCP(*localIP, *portTCP)

	fmt.Println("IP Server: ", ip)

	udpAddr, err := net.ResolveUDPAddr("udp", *localIP + ":" + *portUDP)

	if err != nil {
		log.Fatal(err)
	}

	// connect udp server
	conn, err := net.ListenUDP("udp", udpAddr)

	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}

	//
	for {
		fmt.Printf("Accepting a new packet\n")

		handleUDPConnection(conn)
	}


}

func startServerTCP(ip string, port string) {

	fmt.Println("Starting TCP")
	// Start the TCP server
	serverTCP, err := net.Listen("tcp", fmt.Sprintf("%s:%s", ip, port))

	checkError(err)

	defer serverTCP.Close()

	newConnections := make(chan net.Conn)

	//messages := make(chan string)

	go func() {
		for {
			conn, err := serverTCP.Accept()

			if err != nil {
				fmt.Println("Error: ", err)
				os.Exit(1)
			}

			newConnections <- conn
		}
	}()

	for {
		select {
		case conn := <-newConnections:

			defer conn.Close()

			go func(conn net.Conn) {

				for {
					var req Request

					err := json.NewDecoder(conn).Decode(&req)

					if err != nil {
						break
					}

					res := handleRequest(req)

					json.NewEncoder(conn).Encode(res)
				}


			}(conn)

		}

	}

}


// Handle all connections in UDP Server
func handleUDPConnection(conn *net.UDPConn)  {

	buf := make([]byte, 2048)

	// Read buffer
	n, _, err := conn.ReadFromUDP(buf)

	if err != nil {
		fmt.Printf("net.ReadFromUDP() error: %s\n", err)
		return
	}

	var req Request

	err = json.Unmarshal(buf[:n], &req)

	if err != nil {
		log.Print(err)
		return
	}

	go handleRequest(req)

}


func handleRequest(req Request) Response {

	fmt.Printf("Request: %+v\n", req)

	res := Response{}

	switch req.Action {
	case "GET":

		switch req.Method {

		case "consumptions":
			layoutDate := "02-01-2006"

			startDate, _ := time.Parse(layoutDate, req.Params["startDate"].(string))
			endDate, _ := time.Parse(layoutDate, req.Params["endDate"].(string))

			endDate = endDate.AddDate(0, 0, 1)

			sensorID := int(req.Params["sensorID"].(float64))

			jSensor, err := db.GetSensor(sensorID)

			if err != nil {
				fmt.Println(err)
				res.Error = err.Error()
				break
			}

			measures, err := jSensor.Get("measures").Array()

			if err != nil {
				fmt.Println(err)
				res.Error = err.Error()
				break
			}

			var retAr = make(map[string]interface{})

			// groups all ConsumptionInfo in a specified range
			for _, consumption := range measures {

				c := consumption.(map[string]interface{})

				t, _ := time.Parse(time.RFC3339, c["date"].(string))

				if inTimeSpan(startDate, endDate, t) {
					temp := retAr[t.Format(layoutDate)]

					if temp != nil {
						parse := temp.([]interface{})
						updateList := append(parse, c)
						retAr[t.Format(layoutDate)]= updateList
					} else {
						retAr[t.Format(layoutDate)] = []interface{}{c}
					}
				}

			}

			if len(retAr) > 0 {
				res.Success = true
				res.Data = retAr
			}


		}

	case "POST":

		switch req.Method {
		case "measure":

			id := int(req.Params["id"].(float64))
			date := req.Params["date"].(string)
			measure := uint64(req.Params["measure"].(float64))

			newInfo := ConsumptionInfo{id,date, measure}

			wasSaved := newInfo.save()

			if wasSaved {
				handleMsgRequest(1, reflect.TypeOf(newInfo).Kind().String())

				if measure > 15000 {
					sendMail(id)
				}

				break
			}

			handleMsgRequest(0, reflect.TypeOf(newInfo).Kind().String())

		case "login":
			// Load all users
			j, _ := db.Load("users")

			identify := req.Params["cpf"].(string)

			// Array of users
			data, _ := j.Array()

			found := false

			for i := range data {
				temp := data[i].(map[string]interface{})

				if temp["cpf"].(string) == identify {
					res.Success = true
					res.Data = temp
					found = true
					break
				}
			}

			if !found {
				res.Success = false
				res.Error = "User not found"
			}

		}
	}

	return res
}

func sendMail(id int) {
	jSensor, err := db.GetSensor(id)

	if err != nil {
		fmt.Println(err)
		return
	}

	jUser, err := db.GetUser(jSensor.Get("userID").MustInt())

	if err != nil {
		fmt.Println(err)
		return
	}

	newMail := mail.New(jUser.Get("email").MustString(), "O limite estipulado foi atingindo")

	mail.ConnectAndSend(newMail)
}

func handleMsgRequest(status int, model string) {
	if status == 1 {
		fmt.Println(fmt.Sprintf("Status: OK, Msg: %s saved with success", model))
	} else {
		fmt.Println(fmt.Sprintf("Status: Error, Msg: %s was not saved", model))
	}

}


func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}

//https://play.golang.org/p/BDt3qEQ_2H
func externalIP() (string, error) {
	ifaces, err := net.Interfaces()

	if err != nil {
		return "", err
	}

	for _, iface := range ifaces {

		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}

		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}

		addrs, err := iface.Addrs()

		if err != nil {
			return "", err
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}

	return "", errors.New("are you connected to the network?")
}

// checks if the date is in a certain range
func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}