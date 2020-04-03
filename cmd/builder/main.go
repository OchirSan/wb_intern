package main

import (
	"github.com/OchirSan/wb_intern/tree/Dev/pkg/dish"
	"github.com/OchirSan/wb_intern/tree/Dev/pkg/builder"
	"github.com/OchirSan/wb_intern/tree/Dev/pkg/pilaf"
	"github.com/OchirSan/wb_intern/tree/Dev/pkg/pastaCutlets"
)

func main() {
	dish1 := dish.NewDish()
	dish2 := dish.NewDish()
	a := pastaCutlets.NewPastaCutlets(dish1)
	b := pilaf.NewPilaf(dish2)
	director := builder.NewDirector(a)
	director.BuildDish()
	director1 := builder.NewDirector(b)
	director1.BuildDish()
	a.ShowDish()
	b.ShowDish()
}
