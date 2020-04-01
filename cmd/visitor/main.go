package main

import (
	"fmt"
	"github.com/OchirSan/wb_intern/tree/visitor/pkg/bank"
	"github.com/OchirSan/wb_intern/tree/visitor/pkg/factory"
	"github.com/OchirSan/wb_intern/tree/visitor/pkg/house"
	"github.com/OchirSan/wb_intern/tree/visitor/pkg/street"
	"github.com/OchirSan/wb_intern/tree/visitor/pkg/visitor"
)

func main(){
	house := house.NewHouse("Lenina 23", 120, true)
	factory := factory.NewFactory("Lenina 12", false)
	bank := bank.NewBank("Lenina 117", true)
	street1 := street.NewStreet("Lenina", house, factory, bank)

	visitor1 := visitor.NewVisitor()
	res := street1.Accept(visitor1)
	fmt.Println(res)
}
