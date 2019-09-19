package mhz19

type Mhz19 interface {
	Open() error
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
