package facade

import (
	"testing"
	"https://github.com/OchirSan/wb_intern/tree/master/pkg/carry"
	"https://github.com/OchirSan/wb_intern/tree/master/pkg/support"
	"https://github.com/OchirSan/wb_intern/tree/master/pkg/ruiner"
	"https://github.com/OchirSan/wb_intern/tree/master/pkg/team"
)

var okResult = "команда проигрывает"

func TestPlayTheGame(t *testing.T) {
	support := support.NewSupport("merlin", true)
	carry := carry.NewCarry("artur")
	ruiner := ruiner.NewRuiner("ron")
	team := team.NewTeam(carry, support, ruiner)
	result := team.PlayTheGame(60)
	if result != okResult {
		t.Errorf("Expect result to equal %s, but %s.\n", okResult, result)
	}
}
