package dish

import "fmt"

type Dish interface {
	SetMeat(string)
	SetGarnish(string)
	ShowDish()
}

type dish struct {
	meat    string
	garnish string
}

func (d *dish) SetMeat(s string) {
	d.meat = s
}

func (d *dish) SetGarnish(s string) {
	d.garnish = s
}

// Show Dish
func (d *dish) ShowDish() {
	fmt.Println("Yummi, this", d.meat, "delicious")
	fmt.Println("but garnish", d.garnish, "tasteless")
}

// Create new dish
func NewDish() Dish {
	return &dish{}
}
