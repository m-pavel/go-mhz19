package mhz19

import (
	"io"

	"errors"

	"time"

	"fmt"

	"github.com/jacobsa/go-serial/serial"
)

const (
	readings             = "\xff\x01\x86\x00\x00\x00\x00\x00\x79"
	abc_on               = "\xff\x01\x79\xa0\x00\x00\x00\x00\xe6"
	abc_off              = "\xff\x01\x79\x00\x00\x00\x00\x00\x86"
	zero_point_cal       = "\xff\x01\x87\x00\x00\x00\x00\x00\x78"
	span_point_cal       = "\xff\x01\x88\x00\x00\x00\x00\x00\x00"
	detection_range_5000 = "\xff\x01\x99\x00\x00\x00\x13\x88\xcb"
	detection_range_2000 = "\xff\x01\x99\x00\x00\x00\x07\xd0\x8F"
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

	buffer := make([]uint8, 9)
	time.Sleep(500 * time.Millisecond)
	if n, err = s.port.Read(buffer); err != nil {
		return nil, err
	}

	if n != 9 {
		return nil, errors.New(fmt.Sprintf("Wrong readings (Size %d)", n))
	}

	if buffer[0] == 0xff && buffer[1] == 0x86 {
		return &Readings{
			Co2:         int(buffer[2])<<8 + int(buffer[3]),
			Temperature: int(buffer[4]) - 0x28,
			Tt:          int(buffer[4]),
			Ss:          int(buffer[5]),
			UhUl:        int(buffer[6])<<8 + int(buffer[7]),
		}, nil
	} else {
		return nil, errors.New("Wrong readings (Format)")
	}
}

func (s *serialMhz19) Abc(on bool) error {
	var request []byte
	if on {
		request = []byte(abc_on)
	} else {
		request = []byte(abc_off)
	}

	_, err := s.port.Write(request)
	return err
}

func (s *serialMhz19) SpanPointCalibration(span byte) error {
	_, err := s.port.Write(s.spanRequest(span))
	return err
}

func (s *serialMhz19) spanRequest(span byte) []byte {
	request := []byte(span_point_cal)
	request[3] = span << 8
	request[4] = byte(int(span) % 256)
	request[8] = 0xff - byte(int(0x01+0x88+request[3]+request[4])%0x100) + 1
	return request
}

func (s *serialMhz19) ZeroPointCalibration(span int) error {
	_, err := s.port.Write([]byte(zero_point_cal))
	return err
}

func (s *serialMhz19) DetectionRange5000(span int) error {
	_, err := s.port.Write([]byte(detection_range_5000))
	return err
}

func (s *serialMhz19) DetectionRange2000(span int) error {
	_, err := s.port.Write([]byte(detection_range_2000))
	return err
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
