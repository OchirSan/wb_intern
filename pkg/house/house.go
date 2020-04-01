package house

import (
	"fmt"
)

type houseInterface interface {
	VisitHouse() string
}

type house struct {
	name               string
	FinancialCondition int
	Prudency           bool
}

// Accept visitor in house
func Accept(house houseInterface) string {
	return house.VisitHouse()
}

// Visitor visit house
func (h house) VisitHouse() string {
	res := fmt.Sprintln("visiting", h.name, "house and offer medical insurance")
	if h.Prudency == true && h.FinancialCondition >= 100 {
		res += fmt.Sprintln("accept visitor")
	} else {
		res += fmt.Sprintln("kick off visitor")
	}
	return res
}


// Create new house
func NewHouse(name string, FinancialCondition int, Prudency bool) houseInterface {
	return &house{
		name:               name,
		FinancialCondition: FinancialCondition,
		Prudency:           Prudency,
	}
}
