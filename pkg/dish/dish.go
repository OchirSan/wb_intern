package dish

import "fmt"

type Dish interface {
	SetMeat(string)
	SetGarnish(string)
	ShowDish()
	GetMeat() int
	GetGarnish() int
	IsItDish() string
}

type dish struct {
	meat          string
	garnish       string
	weightMeat    int
	weightGarnish int
}

func (d *dish) SetMeat(s string) {
	d.meat = s
}

func (d *dish) IsItDish() string {
	if d.GetMeat() > 300 || d.GetGarnish() > 400 {
		res := fmt.Sprintln("too much meat or garnish")
		return res
	} else {
		res := fmt.Sprintln("Its, Ok, we can cook it")
		return res
	}
}

func (d *dish) SetGarnish(s string) {
	d.garnish = s
}

func (d *dish) GetMeat() int {
	return d.weightMeat
}

func (d *dish) GetGarnish() int {
	return d.weightGarnish
}

// Show Dish
func (d *dish) ShowDish() {
	fmt.Println("Yummi, this", d.meat, "delicious")
	fmt.Println("but garnish", d.garnish, "tasteless")
}

// Create new dish
func NewDish(weightMeat int, weightGarnish int) Dish {
	return &dish{
		weightMeat:    weightMeat,
		weightGarnish: weightGarnish,
	}
}
