package support

import "fmt"

type supportInterface interface {
	Buff() bool
}

type support struct {
	name      string
	WhoBuffed bool
}

// Support buff random hero
func (s support) Buff() bool {
	buff := s.WhoBuffed
	if buff == false {
		fmt.Println(s.name, "бафает керри")
	} else {
		fmt.Println(s.name, "бафает руинера")
	}
	return buff
}

// Creates a new support
func NewSupport(name string, WhoBuffed bool) supportInterface {
	return &support{
		name:      name,
		WhoBuffed: WhoBuffed,
	}
}
