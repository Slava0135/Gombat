package core

import (
	"gombat/game/util"
)

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

func (gs *GameState) SelectGop(pos util.Position) *Gop {
	for _, team := range gs.Teams {
		for _, gop := range team.Gops {
			if pos.DistanceTo(gop.Pos) < GopSize/2 {
				return gop
			}
		}
	}
	return nil
}

func MoveGop(gop *Gop, pos util.Position) {
	gop.Pos = pos
}
