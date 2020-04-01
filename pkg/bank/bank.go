package bank

import "fmt"

type bankInterface interface {
	VisitBank() string
}

// Accept visitor in bank
func Accept(b bankInterface) string {
	return b.VisitBank()
}

// Visitor visit bank
func (bank bank) VisitBank() string {
	res := fmt.Sprintln("visiting", bank.name, "house and offer robbery insurance")
	if bank.IsHadInsurance {
		res += fmt.Sprintln("accept visitor")
	} else {
		res += fmt.Sprintln("kick off visitor")
	}
	return res
}

type bank struct {
	name           string
	IsHadInsurance bool
}

// Create new bank
func NewBank(name string, IsHadInsurance bool) bankInterface {
	return &bank{
		name:           name,
		IsHadInsurance: IsHadInsurance,
	}
}
