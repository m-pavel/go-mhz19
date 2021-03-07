package mhz19

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	s := serialMhz19{}
	r := s.spanRequest(byte(200))
	fmt.Println(r)
	if r[4] != 0xc8 {
		t.Fail()
	}
	if r[3] != 0x0 {
		t.Fail()
	}
	if r[8] != 0xaf {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	s := serialMhz19{}
	r := s.spanRequest(byte(155))
	fmt.Println(r)
	if r[4] != 0x9b {
		t.Fail()
	}
	if r[3] != 0x0 {
		t.Fail()
	}
	if r[8] != 0xdc {
		t.Fail()
	}
}