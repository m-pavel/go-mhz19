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
	timeout time.Duration
}

func (s *serialS8) Open(timeout time.Duration) error {
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
			ch <- &co2.ReadingsResponse{}
		}
	}()

	var wr *co2.ReadingsResponse

	select {
	case r := <- ch:
		wr = r
	case <-time.After(s.timeout):
		return nil, errors.New("Write timeout")
	}

	if wr.E != nil {
		return nil, wr.E
	}

	time.Sleep(500 * time.Millisecond)

	go func() {
		buffer := make([]uint8, 7)
		if n, err = s.port.Read(buffer); err != nil {
			wr = &co2.ReadingsResponse{E: err}
		} else {
			if n != 7 {
				wr = &co2.ReadingsResponse{E: errors.New(fmt.Sprintf("Wrong readings (Size %d): %v", n, buffer))}
			} else {
				wr = &co2.ReadingsResponse{R: &co2.Readings{
					Co2:         int(buffer[3])<<8 + int(buffer[4]),
				}}
			}
		}
	}()
	select {
	case r := <- ch:
		return r.R, r.E
	case <-time.After(s.timeout):
		return nil, errors.New("Read timeout")
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

