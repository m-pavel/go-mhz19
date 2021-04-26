package co2

import (
	"time"
)

type Device interface {
	Open(timeout time.Duration) error
	Close() error
	Read() (*Readings, error)

}

type Readings struct {
	Co2         int `json:"co2"`
	Temperature int `json:"temperature"`
	Tt          int `json:"tt"`
	Ss          int `json:"ss"`
	UhUl        int `json:"uhul"'`
}


type ReadingsResponse struct {
	R *Readings
	E error
}