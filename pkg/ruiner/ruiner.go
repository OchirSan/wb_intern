package ruiner

import (
	"fmt"
)

type ruinerInterface interface {
	LeaveTheGame()
}

type ruiner struct {
	name string
}

// Ruin the game
func (r ruiner) LeaveTheGame() {
	fmt.Println(r.name, "покидает игру")
}

// Create new ruiner
func NewRuiner(name string) ruinerInterface {
	return &ruiner{
		name: name,
	}
}
