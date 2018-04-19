package main

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"os"
	"net"
	"client-server/customer/config"
	"github.com/therecipe/qt/quickcontrols2"
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"client-server/shared"
	"sync"
	"time"
	"strconv"
)

var conf config.Configuration

type Controller struct {
	// signal envia pra QML
	// slot enviar para GO
	core.QObject

	Conn net.Conn

	_ func(data string) `slot:"login"`
	_ func() 			`slot:"logout"`
	_ func(userID int) string 	`slot:"getConsumptionValue"`
	_ func(userID int, start string, end string) `slot:"getConsumptions"`
	_ func(zone int, start string, end string) `slot:"getConsumptionsZone"`

	_ func(reply string) `signal:"sessionAuthenticated"`
	_ func() 			 `signal:"sessionAuthenticationError"`
	_ func() 			 `signal:"sessionTerminated"`
	_ func() 			 `signal:"sessionLoader"`
	_ func(consumptions string)	`signal:"getConsumption"`
}


type Client struct {
	core.QObject
	_ int `property:"state"`
	_ core.QJsonObject `property:"data"`
}

var controller = NewController(nil)

func main() {

	conf = config.LoadConfig()

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", conf.Host, conf.Port))

	checkError(err)

	controller.Conn = conn

	exit := make(chan bool)

	guiInterface(exit)


	for {
		select {
		case <-exit:
			os.Exit(1)
		}
	}

}

func guiInterface(exit chan bool) {
	core.QCoreApplication_SetApplicationName("Monitoramento de Consumo")
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	quickcontrols2.QQuickStyle_SetStyle("material")

	Client_QmlRegisterType2("Client", 1, 0, "Client")

	var engine = qml.NewQQmlApplicationEngine(nil)
	engine.Load(core.NewQUrl3("qrc:///qml/app.qml", 0))

	engine.RootContext().SetContextProperty("Controller", controller)

	functionsController(exit)

	gui.QGuiApplication_Exec()
}

func functionsController(exit chan bool)  {
	controller.ConnectLogin(func(cpf string) {
		req := shared.Request{"login", "POST", nil}

		req.Params = map[string]interface{}{"cpf": cpf}

		json.NewEncoder(controller.Conn).Encode(req)

		mutex := &sync.Mutex{}

		mutex.Lock()

		controller.SessionLoader()

		mutex.Unlock()

		go func(conn net.Conn) {
			var res shared.Response
			err := json.NewDecoder(conn).Decode(&res)

			if err != nil {
				fmt.Println("Connection Close")
				exit <-true
			}

			if res.Success {
				var user shared.User

				mapstructure.Decode(res.Data, &user)

				d, _ := json.Marshal(user)

				fmt.Println("success login")
				controller.SessionAuthenticated(string(d))
			} else {
				fmt.Println("Failed")
				controller.SessionAuthenticationError()
			}

		}(controller.Conn)

	})

	controller.ConnectLogout(func() {
		controller.SessionTerminated()
	})

	controller.ConnectGetConsumptionValue(func(userID int) string {
		req := shared.Request{"consumptions", "GET", nil}

		t := time.Now().Format("02-01-2006")

		req.Params = map[string]interface{}{"startDate": t, "endDate": t, "sensorID": userID}

		// write in socket
		err := json.NewEncoder(controller.Conn).Encode(req)

		if err != nil {
			return "0"
		}

		var res shared.Response
		// read from socket
		err = json.NewDecoder(controller.Conn).Decode(&res)

		if err != nil || !res.Success {
			return "0"
		}

		// get array of measures
		arr := res.Data[t].([]interface{})

		// last measure at 23:59:59
		last := arr[len(arr) - 1].(map[string]interface{})
		// first measure of the day
		first := arr[0].(map[string]interface{})
		// calculate difference
		result := int(last["measure"].(float64)) - int(first["measure"].(float64))

		return strconv.Itoa(result)

	})

	controller.ConnectGetConsumptions(func(userID int, start string, end string) {
		req := shared.Request{"consumptions", "GET", nil}

		req.Params = map[string]interface{}{"startDate": start, "endDate": end, "sensorID": userID}

		err := json.NewEncoder(controller.Conn).Encode(req)

		if err != nil {
			fmt.Println(err)
			return
		}

		var res shared.Response

		err = json.NewDecoder(controller.Conn).Decode(&res)

		if err != nil || !res.Success {
			fmt.Println(err)
			return
		}

		arr := res.Data

		d, _ := json.Marshal(arr)

		controller.GetConsumption(string(d))
	})

	controller.ConnectGetConsumptionsZone(func(zone int, start string, end string) {
		req := shared.Request{"consumptionsZone", "GET", nil}

		req.Params = map[string]interface{}{"startDate": start, "endDate": end, "zone": zone}

		err := json.NewEncoder(controller.Conn).Encode(req)

		if err != nil {
			fmt.Println(err)
			return
		}

		var res shared.Response

		err = json.NewDecoder(controller.Conn).Decode(&res)

		if err != nil || !res.Success {
			fmt.Println(err)
			return
		}

		arr := res.Data

		d, _ := json.Marshal(arr)

		controller.GetConsumption(string(d))
	})
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}