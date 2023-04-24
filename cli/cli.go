package main

import (
	"fmt"

	co2 "github.com/m-pavel/go-co2/pkg/api"
	"github.com/m-pavel/go-co2/pkg/producer"
	ghm "github.com/m-pavel/go-hassio-mqtt/pkg"
)

func Converter(r *co2.Readings) any {
	return fmt.Sprintf("CO2 %d ppm\nTemp %d C\n", r.Co2, r.Temperature)
}

func main() {
	ghm.NewExecutor[*co2.Readings]("co2", &producer.Co2Service{}, &ghm.ConsoleConsumer[*co2.Readings]{Converter: Converter}).Main()
}
