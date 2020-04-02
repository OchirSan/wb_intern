package visitor_test

import (
	"testing"
	"github.com/OchirSan/wb_intern/tree/visitor/pkg/house"
	"github.com/OchirSan/wb_intern/tree/visitor/pkg/factory"
	"github.com/OchirSan/wb_intern/tree/visitor/pkg/bank"
	"github.com/OchirSan/wb_intern/tree/visitor/pkg/street"
	"github.com/OchirSan/wb_intern/tree/visitor/pkg/visitor"
)

var(
	okresult = "visiting street Lenina street\n" +
		"visiting Lenina 23 house and offer medical insurance\n" +
		"accept visitor\n" +
		"visiting Lenina 117 house and offer robbery insurance\n" +
		"accept visitor\n" +
		"visiting Lenina 12 house and offer fire and flood insurance\n" +
		"kick off visitor"
	streetName = "Lenina street"
	factoryIsHadDanger = false
	factoryName = "Lenina 12"
	bankName = "Lenina 117"
	bankIsHadIndurance = true
	houseName = "Lenina 23"
	houseFinancialCondition = 120
	housePrudency = true
)

func TestStreetAccept(t *testing.T){
	house := house.NewHouse(houseName, houseFinancialCondition, housePrudency)
	factory := factory.NewFactory(factoryName, factoryIsHadDanger)
	bank := bank.NewBank(bankName, bankIsHadIndurance)
	street := street.NewStreet(streetName, house, factory, bank)
	visitor := visitor.NewVisitor()
	res := street.Accept(visitor)
	if res != okresult{
		t.Errorf("Expect result to equal %s, but %s.\n", okresult, res)
	}
}
