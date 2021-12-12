package game

import (
	"Gombat/game/core"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type Game struct {
	Images    *map[string]*ebiten.Image
	GameState *core.GameState
}

func NewGame(images *map[string]*ebiten.Image) *Game {
	g := new(Game)
	g.Images = images
	g.GameState = core.NewGameState(2, 4, 16, 10)
	return g
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
}

func (*Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
