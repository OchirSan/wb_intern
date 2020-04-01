package street

import "fmt"

type Street interface {
	Accept(Visitor) string
}

// Define interface for house, bank, factory
type element interface {
	Accept(Visitor) string
}

type Visitor interface {
	VisitHouse(house) string
	VisitBank(bank) string
	VisitFactory(factory) string
	VisitStreet(street) string
}

type factory struct {
	name        string
	IsHadDanger bool
}

type house struct {
	name               string
	FinancialCondition int
	Prudency           bool
}

type bank struct {
	name           string
	IsHadInsurance bool
}

type street struct {
	name    string
	house   element
	bank    element
	factory element
}

type PrintVisitor struct{}

func (s *house) Accept(visitor Visitor) string {
	return visitor.VisitHouse(*s)
}

// Visitor visit house
func (p *PrintVisitor) VisitHouse(house house) string {
	res := fmt.Sprintln("visiting", house.name, "house and offer medical insurance")
	if house.Prudency == true && house.FinancialCondition >= 100 {
		res += fmt.Sprintln("accept visitor")
	} else {
		res += fmt.Sprintln("kick off visitor")
	}
	return res
}

func (f *factory) Accept(visitor Visitor) string {
	return visitor.VisitFactory(*f)
}

// Visitor visit factory
func (p *PrintVisitor) VisitFactory(factory factory) string {
	res := fmt.Sprintln("visiting", factory.name, "house and offer fire and flood insurance")
	if factory.IsHadDanger {
		res += fmt.Sprintln("accept visitor")
	} else {
		res += fmt.Sprintln("kick off visitor")
	}
	return res
}

func (b *bank) Accept(visitor Visitor) string {
	return visitor.VisitBank(*b)
}

// Visitor visit bank
func (p *PrintVisitor) VisitBank(bank bank) string {
	res := fmt.Sprintln("visiting", bank.name, "house and offer robbery insurance")
	if bank.IsHadInsurance {
		res += fmt.Sprintln("accept visitor")
	} else {
		res += fmt.Sprintln("kick off visitor")
	}
	return res
}

func (s *street) Accept(visitor Visitor) string {
	elements := []element{
		s.house,
		s.bank,
		s.factory,
	}
	res := visitor.VisitStreet(*s)
	// All object on the street accept visitor
	for _, elem := range elements {
		res += elem.Accept(visitor)
	}
	return res
}

func (p *PrintVisitor) VisitStreet(street street) string {
	return fmt.Sprintln("visiting street")
}

// Create new street
func NewStreet(name string, house element, factory element, bank element) Street {
	return &street{
		name:    name,
		house:   house,
		factory: factory,
		bank:    bank,
	}
}

// Create new visitor
func NewVisitor() PrintVisitor {
	return PrintVisitor{}
}

// Create new house
func NewHouse(name string, FinancialCondition int, Prudency bool) element {
	return &house{
		name:               name,
		FinancialCondition: FinancialCondition,
		Prudency:           Prudency,
	}
}

// Create new bank
func NewBank(name string, IsHadInsurance bool) element {
	return &bank{
		name:           name,
		IsHadInsurance: IsHadInsurance,
	}
}

// Create new factory
func NewFactory(name string, IsHadDanger bool) element {
	return &factory{
		name:        name,
		IsHadDanger: IsHadDanger,
	}
}
