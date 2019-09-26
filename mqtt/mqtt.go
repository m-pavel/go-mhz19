package main

import (
	"flag"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/m-pavel/go-hassio-mqtt/pkg"
	"github.com/m-pavel/go-mhz19/pkg"
)

type Co2Service struct {
	m      mhz19.Mhz19
	device *string
}

func (ts *Co2Service) PrepareCommandLineParams() {
	ts.device = flag.String("device", "/dev/serial0", "Serial device")
}
func (ts Co2Service) Name() string { return "mhz19" }

func (ts *Co2Service) Init(client MQTT.Client, topic, topicc, topica string, debug bool, ss ghm.SendState) error {
	ts.m = mhz19.NewSerial(*ts.device)
	return ts.m.Open()
}

func (ts Co2Service) Do() (interface{}, error) {
	return ts.m.Read()
}

func (ts Co2Service) Close() error {
	return ts.m.Close()
}

func main() {
	ghm.NewStub(&Co2Service{}).Main()
}
