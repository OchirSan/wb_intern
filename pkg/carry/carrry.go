package carry

import (
	"fmt"
)

type carryInterface interface {
	KillEnemy(string);
}

type carry struct{
	name string;
	IsBuffed bool;
}

//Win the game
func (c carry) KillEnemy(name string) {
	fmt.Println(name, "убивает врага")
}

//Creates new carry
func NewCarry() carryInterface {
	return &carry{}
}
