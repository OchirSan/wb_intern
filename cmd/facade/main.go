package main

import (
	"github.com/OchirSan/wb_intern/pkg/carry"
	"github.com/OchirSan/wb_intern/pkg/support"
	"github.com/OchirSan/wb_intern/pkg/ruiner"
	"github.com/OchirSan/wb_intern/pkg/team"
)

func main(){
	carry:= carry.NewCarry("artur")
	support := support.NewSupport("merlin", true)
	ruiner := ruiner.NewRuiner("ron")
	team := team.NewTeam(carry, support, ruiner)
	team.PlayTheGame(60)
}
