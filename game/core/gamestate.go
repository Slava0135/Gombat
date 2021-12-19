package core

import (
	"gombat/game/util"
	"math"
)

type GameState struct {
	World *World
	Teams []*Team
}

type Gop struct {
	Health int
	Team   *Team
	Pos    util.Vec2
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
			t.Gops[j] = &Gop{Health: 3, Team: t, Pos: util.Vec2{X: float64(4 * i), Y: float64(4 * j)}}
		}
		gs.Teams[i] = t
	}
	return gs
}

func (gs *GameState) SelectGop(pos util.Vec2) *Gop {
	for _, team := range gs.Teams {
		for _, gop := range team.Gops {
			if pos.DistanceTo(gop.Pos) < GopSize/2 {
				return gop
			}
		}
	}
	return nil
}

func (g *Gop) MoveGop(pos util.Vec2) {
	g.Pos = pos
}

func (w *World) CanMoveGop(from, to util.Vec2) bool {

	grid := make([][]bool, w.Width)
	for i := range grid {
		grid[i] = make([]bool, w.Height)
	}
	for i := 0; i < w.Width; i++ {
		for j := 0; j < w.Height; j++ {
			grid[i][j] = w.Floors[i][j].Passable == false || w.Blocks[i][j].Solid == true
		}
	}

	if _, _, collided := util.RayTrace(grid, from, to); collided {
		return false
	}

	vector := util.Vec2{to.X - from.X, to.Y - from.Y}
	vector = vector.Rotate(math.Pi / 2)

	dx := GopSize / 2 * vector.X / vector.Norm()
	dy := GopSize / 2 * vector.Y / vector.Norm()

	offset1 := util.Vec2{dx, dy}
	offset2 := util.Vec2{-dx, -dy}

	if _, _, collided := util.RayTrace(grid, from.Add(offset1), to.Add(offset1)); collided {
		return false
	}

	if _, _, collided := util.RayTrace(grid, from.Add(offset2), to.Add(offset2)); collided {
		return false
	}

	return true
}
