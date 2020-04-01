package visitor

import (
	"fmt"
	"github.com/OchirSan/wb_intern/tree/visitor/pkg/street"
)

type Visitor interface {
	VisitHouse(street.HouseInterface) string
	VisitBank(street.BankInterface) string
	VisitFactory(street.FactoryInterface) string
	VisitStreet(street.StreetInterface) string
}

type visitor struct{}

// Visitor visit house
func (v visitor) VisitHouse(h street.HouseInterface) string {
	res := fmt.Sprintln("visiting", h.GetName(), "house and offer medical insurance")
	if h.GetPrudency() == true && h.GetFinancialCondition() >= 100 {
		res += fmt.Sprintln("accept visitor")
	} else {
		res += fmt.Sprintln("kick off visitor")
	}
	return res
}

// Visitor visit bank
func (v visitor) VisitBank(bank street.BankInterface) string {
	res := fmt.Sprintln("visiting", bank.GetName(), "house and offer robbery insurance")
	if bank.GetIsHadIndurance() {
		res += fmt.Sprintln("accept visitor")
	} else {
		res += fmt.Sprintln("kick off visitor")
	}
	return res
}

// Visitor visit factory
func (v visitor) VisitFactory(factory street.FactoryInterface) string {
	res := fmt.Sprintln("visiting", factory.GetName(), "house and offer fire and flood insurance")
	if factory.GetIsHadDanger() {
		res += fmt.Sprintln("accept visitor")
	} else {
		res += fmt.Sprintln("kick off visitor")
	}
	return res
}

// Visitor visit street
func (v visitor) VisitStreet(s street.StreetInterface) string {
	res := fmt.Sprintln("visiting street", s.GetName())
	res += v.VisitHouse(s.GetHouse())
	res += v.VisitBank(s.GetBank())
	res += v.VisitFactory(s.GetFactory())
	return res
}

// Create new visitor
func NewVisitor() Visitor {
	return &visitor{}
}
