package s8

import (
	"errors"
	"fmt"
	"time"

	co2 "github.com/m-pavel/go-co2/pkg"
	"github.com/tarm/serial"
)

const (
	readings       = "\xFE\x04\x00\x00\x00\x04\xE5\xC6"
	default_device = "/dev/serial0"
)

type serialS8 struct {
	dev     string
	port    *serial.Port
	timeout time.Duration
}

func (s *serialS8) Open(timeout time.Duration) error {
	var err error
	c := &serial.Config{Name: s.dev, Baud: 9600}
	s.port, err = serial.OpenPort(c)
	s.timeout = timeout
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
	ch := make(chan *co2.ReadingsResponse)

	go func() {
		if n, err = s.port.Write([]byte(readings)); err != nil {
			ch <- &co2.ReadingsResponse{E: err}
		} else {
			if n != 8 {
				ch <- &co2.ReadingsResponse{E: errors.New(fmt.Sprintf("Wrong writings : %d", n))}
			} else {
				ch <- &co2.ReadingsResponse{}
			}
		}
	}()

	{
		var wr *co2.ReadingsResponse
		select {
		case r := <-ch:
			wr = r
		case <-time.After(s.timeout):
			return nil, errors.New("write timeout")
		}

		if wr.E != nil {
			return nil, wr.E
		}
	}

	time.Sleep(300 * time.Millisecond)

	go func() {
		sz := 13
		buffer := make([]byte, sz)
		if n, err = s.port.Read(buffer); err != nil {
			ch <- &co2.ReadingsResponse{E: err}
		} else {
			if n != sz {
				ch <- &co2.ReadingsResponse{E: errors.New(fmt.Sprintf("Wrong readings (Size %d): %v", n, buffer))}
			} else {
				length := uint(buffer[2])
				//status := (uint(buffer[3]) << 8) | uint(buffer[4])
				ppm := (uint(buffer[length+1]) << 8) | uint(buffer[length+2])
				ch <- &co2.ReadingsResponse{R: &co2.Readings{
					Co2: int(ppm),
				}}
			}
		}
	}()
	select {
	case r := <-ch:
		return r.R, r.E
	case <-time.After(s.timeout):
		return nil, errors.New("read timeout")
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
