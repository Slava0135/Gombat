package game

import (
	"Gombat/game/assets"
	"Gombat/game/core"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type Game struct {
	Images    *assets.Images
	GameState *core.GameState
}

func NewGame(images *assets.Images) *Game {
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
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(10, 10)
	op.GeoM.Translate(40, 40)
	screen.DrawImage(g.Images.Gop, op)
	op.GeoM.Translate(40, 40)
	screen.DrawImage(g.Images.Glass, op)
	op.GeoM.Translate(40, 40)
	screen.DrawImage(g.Images.Planks, op)
	op.GeoM.Translate(40, 40)
	screen.DrawImage(g.Images.Concrete, op)
	op.GeoM.Translate(40, 40)
	screen.DrawImage(g.Images.Bricks, op)
	op.GeoM.Translate(40, 40)
	screen.DrawImage(g.Images.Road, op)
	op.GeoM.Translate(40, 40)
	screen.DrawImage(g.Images.Water, op)
	op.GeoM.Translate(40, 40)
	screen.DrawImage(g.Images.Grass, op)
	op.GeoM.Translate(40, 40)
}

func (*Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
