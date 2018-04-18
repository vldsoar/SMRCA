package main

import (
	"net"
	"time"
	"fmt"
	"log"
	"os"
	"encoding/json"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quickcontrols2"
	"github.com/therecipe/qt/quick"
	"sync/atomic"
	"io/ioutil"

	"path"

	"github.com/bitly/go-simplejson"
	"strconv"
	"client-server/shared"
)

type SensorConf struct {
	ID		int `json:"id"`
	Address string `json:"address"`
	InitCounter uint64 `json:"counter"`
	Counter *counter
}

type Controller struct {
	core.QObject
	ConnUDP UDPServ

	_ func(data string) `signal:"updateCounter"`
	_ func(data int) 	`slot:"updateFlowRate"`
}

type UDPServ struct {
	conn *net.UDPConn
}

type counter uint64

func (s *SensorConf) load() {
	dir, err := os.Getwd()

	f, err := ioutil.ReadFile(path.Join(dir, "config.json"))

	println(path.Join(dir, "config.json"))

	if err != nil {
		println(f)
		println(err)
		return
	}

	//checkError(err)
	//
	j, err := simplejson.NewJson(f)

	if err != nil {
		println(err)
		return
	}

	s.Address, _ 		= j.Get("address").String()
	s.ID, _	 		= j.Get("id").Int()
	s.InitCounter, _	= j.Get("counter").Uint64()
}

func (s *SensorConf) save() {

	s.InitCounter = s.Counter.get()

	sensorJson, _ := json.Marshal(s)

	_ = ioutil.WriteFile("config.json", sensorJson, 0644)
}

func (udp *UDPServ) sendToServer(req shared.Request) {
	jsonRequest, err := json.Marshal(req)

	if err != nil { log.Print("Error: ", err) }

	udp.conn.Write(jsonRequest)
}

func (s *UDPServ) start(addr string)  {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)

	if err != nil {
		log.Print("Resolve server address failed.")
		log.Fatal(err)
	}

	s.conn, err = net.DialUDP("udp", nil,  udpAddr)

	if err != nil {
		log.Print("Listen UDP failed.")
		log.Fatal(err)
	}

}

func (c *counter) increment(numberToInc int) uint64 {
	var next uint64

	for {
		next = uint64(*c) + uint64(numberToInc)
		if atomic.CompareAndSwapUint64((*uint64)(c), uint64(*c), next) {
			return next
		}
	}
}

func (c *counter) get() uint64 {
	return atomic.LoadUint64((*uint64)(c))
}


func run() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	quickcontrols2.QQuickStyle_SetStyle("material")


	var view = quick.NewQQuickView(nil)
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)

	var appSensor SensorConf
	appSensor.load()

	count := counter(appSensor.InitCounter)
	appSensor.Counter = &count

	fmt.Println(appSensor)

	flowRateSignal := make(chan int)

	var controller = NewController(nil)

	controller.ConnUDP.start(appSensor.Address)
	// each time you change the value of the water flow
	// send 'signal'
	controller.ConnectUpdateFlowRate(func(data int) {
		go func() {
			flowRateSignal <-data
		}()

	})

	view.RootContext().SetContextProperty("Controller", controller)

	view.SetSource(core.NewQUrl3("qrc:///qml/app.qml", 0))

	go func(control *Controller, count *counter) {
		var flowRate int = 0
		for {
			select {
			case fr := <-flowRateSignal: //receive signal of change
				flowRate = fr
			case <-time.After(time.Second * 1):

				if flowRate == 0 { break }

				control.UpdateCounter(strconv.FormatUint(count.increment(flowRate), 10))
			}
		}
	}(controller, &count)

	go handleSignal(controller, &appSensor)

	//view.Show()

	gui.QGuiApplication_Exec()
}

func main() {

	run()
}

func handleSignal(control *Controller, sensor *SensorConf) {


	for {
		select {

		case <-time.After(time.Second * 7):
			params := make(map[string]interface{})

			params["id"] = sensor.ID
			params["date"] = time.Now()

			sensor.save()

			params["measure"] = sensor.InitCounter

			req := shared.Request{"measure", "POST", params}

			control.ConnUDP.sendToServer(req)

		}
	}
}


func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(0)
	}
}
