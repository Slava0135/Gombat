package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gombat/game/assets"
	"gombat/game/core"
	"gombat/game/view"
)

type Game struct {
	GameState *core.GameState
}

func NewGame() *Game {
	g := new(Game)
	g.GameState = core.NewGameState(2, 4)
	m, _ := assets.LoadMap("test")
	g.GameState.World = m
	return g
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	view.DrawGameState(screen, g.GameState)
}

func (*Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
