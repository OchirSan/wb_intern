package main

import (
	"fmt"
	"github.com/OchirSan/wb_intern/tree/visitor/pkg/bank"
	"github.com/OchirSan/wb_intern/tree/visitor/pkg/factory"
	"github.com/OchirSan/wb_intern/tree/visitor/pkg/house"
	"github.com/OchirSan/wb_intern/tree/visitor/pkg/street"
)

func main(){
	house:= house.NewHouse("Lenina 23", 120, true)
	factory:= factory.NewFactory("Lenina 12", false)
	bank:= bank.NewBank("Lenina 117", true)
	street1 := street.NewStreet("Lenina", house, factory, bank)

	visitor := street.NewVisitor()
	res := visitor.Accept(street1)
	fmt.Println(res)
}
