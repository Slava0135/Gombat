package core

import "gombat/game/util"

type GameState struct {
	World *World
	Teams []*Team
}

type Gop struct {
	Health int
	Team   *Team
	Pos    util.Position
}

type Team struct {
	Id   int
	Gops []*Gop
}

func NewGameState(teamAmount, teamSize int) *GameState {
	gs := new(GameState)
	gs.Teams = make([]*Team, teamAmount)
	for i := range gs.Teams {
		t := &Team{i, make([]*Gop, teamSize)}
		for j := range t.Gops {
			t.Gops[j] = &Gop{Health: 3, Team: t, Pos: util.Position{X: float64(4 * i), Y: float64(4 * j)}}
		}
		gs.Teams[i] = t
	}
	return gs
}
