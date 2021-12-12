package core

type GameState struct {
	World *World
	Teams []*Team
}

type Gop struct {
	Health int
	Team   *Team
	Pos    Position
}

type Position struct {
	X, Y int
}

type Team struct {
	Gops []*Gop
}

func NewGameState(teamAmount, teamSize int) *GameState {
	gs := new(GameState)
	gs.Teams = make([]*Team, teamAmount)
	for i := range gs.Teams {
		t := new(Team)
		t.Gops = make([]*Gop, teamSize)
		for j := range t.Gops {
			t.Gops[j] = &Gop{3, t, Position{0, 0}}
		}
		gs.Teams[i] = t
	}
	return gs
}
