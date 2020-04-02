package carry

import (
	"fmt"
)

type carryInterface interface {
	KillEnemy();
}

type carry struct{
	name string;
}

// Win the game
func (c carry) KillEnemy() {
	fmt.Println(c.name, "убивает врага")
}

// Creates new carry
func NewCarry(name string) carryInterface {
	return &carry{
		name: 	name,
	}
}
