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
	fmt.Println(r)
	fmt.Println(err)

}
