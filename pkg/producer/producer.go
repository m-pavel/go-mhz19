package producer

import (
	"time"

	"github.com/m-pavel/go-co2/pkg/mhz19"
	"github.com/m-pavel/go-co2/pkg/s8"
	"github.com/spf13/cobra"

	co2 "github.com/m-pavel/go-co2/pkg/api"
)

type Co2Service struct {
	d      co2.Device
	device string
	dtype  string
}

func (ts *Co2Service) Setup(cmd *cobra.Command, name string) {
	cmd.PersistentFlags().StringVar(&ts.device, "device", "/dev/serial0", "Serial device")
	cmd.PersistentFlags().StringVar(&ts.dtype, "type", "mhz19", "mhz19 or s8")
}
func (ts *Co2Service) Init(bool) error {
	switch ts.dtype {
	case "mhz19":
		ts.d = mhz19.NewSerial(ts.device)
	case "s8":
		ts.d = s8.NewSerial(ts.device)
	default:
		panic("Wrong device type " + ts.dtype)
	}
	return ts.d.Open(time.Second * 5)
}

func (ts *Co2Service) Produce() (*co2.Readings, error) {
	return ts.d.Read()
}

func (ts *Co2Service) Close() error {
	return ts.d.Close()
}
