package mhz19

import (
	"io"

	"errors"

	"fmt"

	"github.com/jacobsa/go-serial/serial"
)

const (
	readings = "\xff\x01\x86\x00\x00\x00\x00\x00\x79"
)

type serialMhz19 struct {
	dev  string
	port io.ReadWriteCloser
}

func (s *serialMhz19) Open() error {
	options := serial.OpenOptions{
		PortName:        s.dev,
		BaudRate:        9600,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
		ParityMode:      serial.PARITY_NONE,
	}

	var err error
	s.port, err = serial.Open(options)
	return err
}

func (s *serialMhz19) Close() error {
	var err error
	if s.port != nil {
		err = s.port.Close()
		s.port = nil
	}
	return err
}
func (s *serialMhz19) Read() (*Readings, error) {
	var n int
	var err error
	if n, err = s.port.Write([]byte(readings)); err != nil {
		return nil, err
	}
	buffer := make([]byte, 9)

	if n, err = s.port.Read(buffer); err != nil {
		return nil, err
	}

	fmt.Println(n)
	fmt.Println(buffer)
	if n != 9 {
		return nil, errors.New("Wrong readings")
	}

	if buffer[0] == '\xff' && buffer[1] == '\x86' {
		return &Readings{
			Co2:         int(buffer[2])*256 + int(buffer[3]),
			Temperature: int(buffer[4]) - 40,
			Tt:          int(buffer[4]),
			Ss:          int(buffer[5]),
			UhUl:        int(buffer[6])*256 + int(buffer[7]),
		}, nil
	} else {
		return nil, errors.New("Wrong readings")
	}
}

func NewSerial(device ...string) Mhz19 {
	sm := serialMhz19{}
	if len(device) == 1 {
		sm.dev = device[0]
	} else {
		sm.dev = "/dev/serial0"
	}
	return &sm
}
