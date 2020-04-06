package main

import (
	"github.com/OchirSan/wb_intern/tree/Dev/pkg/dish"
	"github.com/OchirSan/wb_intern/tree/Dev/pkg/builder"
	"github.com/OchirSan/wb_intern/tree/Dev/pkg/pilaf"
	"github.com/OchirSan/wb_intern/tree/Dev/pkg/pastaCutlets"
)

func main() {
	result := "too much meat or garnish"
	dish1 := dish.NewDish(200, 150)
	dish2 := dish.NewDish(250, 200)
	res1 := dish1.IsItDish()
	res2 := dish2.IsItDish()
	if res1 != result {
		a := pasta_cutlets.NewPastaCutlets(dish1)
		director := builder.NewDirector(a)
		director.BuildDish()
		a.ShowDish()

	}
	if res2 != result {
		b := pilaf.NewPilaf(dish2)
		director1 := builder.NewDirector(b)
		director1.BuildDish()
		b.ShowDish()
	}
}
