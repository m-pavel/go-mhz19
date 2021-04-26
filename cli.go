package main

import (
	"fmt"
	co22 "github.com/m-pavel/go-co2/pkg"
	"github.com/m-pavel/go-co2/pkg/s8"

	"flag"

	"log"

	"github.com/m-pavel/go-co2/pkg/mhz19"
)

func main() {
	device := flag.String("device", "/dev/serial0", "Serial device")
	dtype := flag.String("type", "mhz19", "mhz19 or s8")
	flag.Parse()

	var co2d co22.Device
	switch *dtype {
	case "mhz19":
		co2d = mhz19.NewSerial(*device)
	case "s8":
		co2d = s8.NewSerial(*device)
	default:
		panic("Wrong device type " + *dtype)
	}

	if err := co2d.Open(); err != nil {
		log.Panic(err)
	}
	log.Printf("Opened %s", *device)
	defer co2d.Close()
	r, err := co2d.Read()
	if err == nil {
		fmt.Printf("CO2 %d ppm\nTemp %d C\n", r.Co2, r.Temperature)
	} else {
		fmt.Println(err)
	}

}
