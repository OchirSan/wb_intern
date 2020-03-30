package support

import (
	"fmt"
	"math/rand"
	"time"
)

type supportInterface interface {
	Buff(string) int;
}

type support struct{
	name string;
}

//Support buff random hero
func (s support) Buff(name string) int {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(name, "бафает случайного героя")
	return rand.Intn(2);
}

//Creates a new support
func NewSupport() supportInterface {
	return &support{}
}
