package ruiner

import (
	"fmt"
)

type ruinerInterface interface {
	LeaveTheGame(string);
}

type ruiner struct{
	name string;
}

//Ruin the game
func (r ruiner) LeaveTheGame(name string){
	fmt.Println(name, "покидает игру")
}

//Create new ruiner
func NewRuiner() ruinerInterface {
	return &ruiner{}
}
