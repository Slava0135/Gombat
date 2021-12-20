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
			if pos.IsInSquareBounds(gop.Pos, GopSize/2) {
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

	collisionGrid := make([][]bool, w.Width)
	for i := range collisionGrid {
		collisionGrid[i] = make([]bool, w.Height)
	}
	for i := 0; i < w.Width; i++ {
		for j := 0; j < w.Height; j++ {
			collisionGrid[i][j] = w.Floors[i][j].Passable == false || w.Blocks[i][j].Solid == true
		}
	}

	{
		if _, _, collided := util.RayTrace(collisionGrid, from, to); collided {
			return false
		}

		dx := GopSize / math.Sqrt(2) / 2
		dy := GopSize / math.Sqrt(2) / 2

		offset := util.Vec2{dx, dy}
		if _, _, collided := util.RayTrace(collisionGrid, from.Add(offset), to.Add(offset)); collided {
			return false
		}

		offset = util.Vec2{-dx, -dy}
		if _, _, collided := util.RayTrace(collisionGrid, from.Add(offset), to.Add(offset)); collided {
			return false
		}

		offset = util.Vec2{-dx, dy}
		if _, _, collided := util.RayTrace(collisionGrid, from.Add(offset), to.Add(offset)); collided {
			return false
		}

		offset = util.Vec2{dx, -dy}
		if _, _, collided := util.RayTrace(collisionGrid, from.Add(offset), to.Add(offset)); collided {
			return false
		}
	}

	for x := 0; x < w.Width; x++ {
		for y := 0; y < w.Height; y++ {
			if collisionGrid[x][y] {
				if to.IsInSquareBounds(util.Vec2{float64(x) + 0.5, float64(y) + 0.5}, 0.5+GopSize/2) {
					return false
				}
			}
		}
	}
	return true
}

func (gs *GameState) Shoot(from, to util.Vec2) {
	w := gs.World
	collisionGrid := make([][]bool, w.Width)
	for i := range collisionGrid {
		collisionGrid[i] = make([]bool, w.Height)
	}
	for i := 0; i < w.Width; i++ {
		for j := 0; j < w.Height; j++ {
			collisionGrid[i][j] = w.Blocks[i][j].Solid == true
		}
	}

	x, y, _ := util.RayTraceUntilHit(collisionGrid, from, to)
	blockHit := util.Vec2{float64(x), float64(y)}
	hitDist := from.DistanceTo(blockHit)

	minGopDist := math.Inf(1)
	for _, t := range gs.Teams {
		for _, g := range t.Gops {
			if g.Pos != from && util.DoesLineIntersectSquare(from, to, g.Pos, GopSize) {
				dist := from.DistanceTo(g.Pos)
				if blockHit.DistanceTo(g.Pos)+GopSize/2 < dist+hitDist { // gop is in correct direction
					minGopDist = math.Min(dist, minGopDist)
				}
			}
		}
	}

	if minGopDist < hitDist {

	} else if util.IsInBounds(x, y, w.Width, w.Height) {
		if w.Blocks[x][y].Destructible {
			w.Blocks[x][y] = Blocks[0]
		}
	}
}
