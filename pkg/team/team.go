package team

import "fmt"

type carryInterface interface {
	KillEnemy(string);
}

type supportInterface interface {
	Buff(string) int;
}

type ruinerInterface interface {
	LeaveTheGame(string);
}

type Team interface {
	PlayTheGame()
}

type team struct {
	player1 carryInterface
	player2 supportInterface
	player3 ruinerInterface
}

// Describe the game
func (t *team) PlayTheGame() {
	//Define 0 or 1 random, 0 is a carry, 1 is a ruiner
	var a int = t.player2.Buff("merlin")
	if a == 0 {
		t.player1.KillEnemy("artur")
		fmt.Println("команда выигрывает")
	} else {
		t.player3.LeaveTheGame("ron")
		fmt.Println("команда проигрывает")
	}
}

//Creates new team
func NewTeam(player1 carryInterface, player2 supportInterface, player3 ruinerInterface) Team {
	return &team{
		player1:   player1,
		player2:   player2,
		player3:   player3,
	}
}
