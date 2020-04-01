package team

import "fmt"

type carryInterface interface {
	KillEnemy()
}

type supportInterface interface {
	Buff() bool
}

type ruinerInterface interface {
	LeaveTheGame()
}

type Team interface {
	PlayTheGame(moral int)
}

type team struct {
	player1        carryInterface
	player2        supportInterface
	player3        ruinerInterface
	moralCondition int
}

// Describe the game
func (t *team) PlayTheGame(moral int) {
	t.moralCondition = moral
	// Define false or true, false is a carry, true is a ruiner
	var a bool = t.player2.Buff()
	if a == false {
		t.moralCondition += 50
	} else {
		t.moralCondition -= 50
	}
	if t.moralCondition >= 100 {
		t.player1.KillEnemy()
		fmt.Println("команда выигрывает")
	} else {
		t.player3.LeaveTheGame()
		fmt.Println("команда проигрывает")
	}
}

// Creates new team
func NewTeam(player1 carryInterface, player2 supportInterface, player3 ruinerInterface) Team {
	return &team{
		player1:        player1,
		player2:        player2,
		player3:        player3,
		moralCondition: 0,
	}
}
