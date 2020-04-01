package street

import "fmt"

type Visitor interface {
	Accept(streetInterface) string
}

type bankInterface interface {
	VisitBank() string
}

type houseInterface interface {
	VisitHouse() string
}

type factoryInterface interface {
	VisitFactory() string
}

type streetInterface interface {
	VisitStreet() string
}

type street struct {
	name    string
	house   houseInterface
	bank    bankInterface
	factory factoryInterface
}

type visitor struct{}

func (v *visitor) Accept(s streetInterface) string {
	return s.VisitStreet()
}

// Visitor visit street
func (s street) VisitStreet() string {
	res := fmt.Sprintln("visiting street", s.name)
	res += s.house.VisitHouse()
	res += s.bank.VisitBank()
	res += s.factory.VisitFactory()
	return res
}

// Create new street
func NewStreet(name string, house houseInterface, factory factoryInterface, bank bankInterface) streetInterface {
	return &street{
		name:    name,
		house:   house,
		factory: factory,
		bank:    bank,
	}
}

// Create new visitor
func NewVisitor() Visitor {
	return &visitor{}
}
