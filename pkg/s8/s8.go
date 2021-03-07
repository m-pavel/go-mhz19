package s8

import (
	"errors"
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	co2 "github.com/m-pavel/go-co2/pkg"
	"io"
	"time"
)

const (
	readings             = "\xFE\x44\x00\x08\x02\x9F\x25"
	default_device = "/dev/serial0"
)

type serialS8 struct {
	dev  string
	port io.ReadWriteCloser
}

func (s *serialS8) Open() error {
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

func (s *serialS8) Close() error {
	var err error
	if s.port != nil {
		err = s.port.Close()
		s.port = nil
	}
	return err
}
func (s *serialS8) Read() (*co2.Readings, error) {
	var n int
	var err error
	if n, err = s.port.Write([]byte(readings)); err != nil {
		return nil, err
	}

	buffer := make([]uint8, 9)
	time.Sleep(500 * time.Millisecond)
	if n, err = s.port.Read(buffer); err != nil {
		return nil, err
	}

	if n != 9 {
		return nil, errors.New(fmt.Sprintf("Wrong readings (Size %d): %v", n, buffer))
	}

	if buffer[0] == 0xff && buffer[1] == 0x86 {
		return &co2.Readings{
			Co2:         int(buffer[2])<<8 + int(buffer[3]),
			Temperature: int(buffer[4]) - 0x28,
			Tt:          int(buffer[4]),
			Ss:          int(buffer[5]),
			UhUl:        int(buffer[6])<<8 + int(buffer[7]),
		}, nil
	} else {
		return nil, errors.New(fmt.Sprintf("Wrong readings (Format): %v", buffer))
	}
}

func NewSerial(device ...string) co2.Device {
	sm := serialS8{}
	if len(device) == 1 {
		sm.dev = device[0]
	} else {
		sm.dev = default_device
	}
	return &sm
}

