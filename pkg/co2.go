package co2

type Device interface {
	Open() error
	Close() error
	Read() (*Readings, error)
	Abc(bool) error
}

type Readings struct {
	Co2         int `json:"co2"`
	Temperature int `json:"temperature"`
	Tt          int `json:"tt"`
	Ss          int `json:"ss"`
	UhUl        int `json:"uhul"'`
}
