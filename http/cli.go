package main

import (
	co2 "github.com/m-pavel/go-co2/pkg/api"
	"github.com/m-pavel/go-co2/pkg/producer"
	ghm "github.com/m-pavel/go-hassio-mqtt/pkg"
)

func Converter(r *co2.Readings) ghm.Entry {
	//return ghm.Entry{"co2": r.Co2, "temp": r.Temperature}
	return ghm.Entry{"value": r.Co2}
}

func main() {
	ghm.NewExecutor[*co2.Readings]("co2", &producer.Co2Service{}, &ghm.HttpServer[*co2.Readings]{ToRawConverter: Converter}).Main()
}
