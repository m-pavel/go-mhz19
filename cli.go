package main

import (
	"fmt"

	"github.com/m-pavel/go-mhz19/pkg"
)

func main() {
	mhz := mhz19.NewSerial()
	if err := mhz.Open(); err != nil {
		panic(err)
	}
	defer mhz.Close()
	r, err := mhz.Read()
	if err == nil {
		fmt.Printf("CO2 %d ppm\nTemp %d C\n", r.Co2, r.Temperature)
	} else {
		fmt.Println(err)
	}

}
