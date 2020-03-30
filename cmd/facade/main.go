package main

import (
	"pkg/carry"
	"pkg/ruiner"
	"pkg/support"
	"pkg/team"
)

func main(){
	T := team.NewTeam(carry.NewCarry(), support.NewSupport(), ruiner.NewRuiner())
	T.PlayTheGame()
}
