package builder_test

import (
	"github.com/OchirSan/wb_intern/tree/Dev/pkg/dish"
	"testing"
)

var (
	okresult1      = "too much meat or garnish"
	okresult2      = "Its, Ok, we can cook it"
	s              string
	weightMeat2    = 500
	weightGarnish2 = 100
)

func TestIsItDish(t *testing.T) {
	dish := dish.NewDish(weightMeat2, weightGarnish2)
	s := dish.IsItDish()
	if dish.GetMeat() > 300 || dish.GetGarnish() > 400 {
		if s != okresult1 {
			t.Errorf("Expect result to equal %s, but %s.\n", okresult1, s)
		}
	} else {
		if s != okresult2 {
			t.Errorf("Expect result to equal %s, but %s.\n", okresult2, s)
		}
	}
}
