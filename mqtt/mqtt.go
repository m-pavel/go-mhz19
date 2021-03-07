package main

import (
	"flag"
	co2 "github.com/m-pavel/go-co2/pkg"
	"github.com/m-pavel/go-co2/pkg/s8"

	"github.com/m-pavel/go-hassio-mqtt/pkg"
	"github.com/m-pavel/go-co2/pkg/mhz19"
)

type Co2Service struct {
	ghm.NonListerningService
	d      co2.Device
	device *string
	dtype *string
}

func (ts *Co2Service) PrepareCommandLineParams() {
	ts.device = flag.String("device", "/dev/serial0", "Serial device")
	ts.dtype = flag.String("type", "mzh19", "mzh19 or s8")
}
func (ts Co2Service) Name() string { return "co2" }

func (ts *Co2Service) Init(ctx *ghm.ServiceContext) error {
	switch *ts.dtype {
	case "mhz19":
		ts.d = mhz19.NewSerial(*ts.device)
	case "s8":
		ts.d = s8.NewSerial(*ts.device)
	default:
		panic("Wrong device type" + *ts.dtype)
	}
	return ts.d.Open()
}

func (ts Co2Service) Do() (interface{}, error) {
	return ts.d.Read()
}

func (ts Co2Service) Close() error {
	return ts.d.Close()
}

func main() {
	ghm.NewStub(&Co2Service{}).Main()
}
