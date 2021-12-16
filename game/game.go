package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gombat/game/assets"
	"gombat/game/core"
	"gombat/game/util"
	"gombat/game/view"
	"math"
)

const scrollSpeed = 32

type Game struct {
	GameState   *core.GameState
	ViewOptions *view.ViewOptions
}

func NewGame() *Game {
	g := new(Game)
	g.GameState = core.NewGameState(2, 4)
	m, _ := assets.LoadMap("test3")
	g.GameState.World = m
	g.ViewOptions = &view.ViewOptions{
		CameraPos: &util.Position{},
		Scale:     4,
	}
	return g
}

func (g *Game) Update() error {
	g.UpdateControls()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	view.DrawGameState(screen, g.ViewOptions, g.GameState)
}

func (*Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *Game) UpdateControls() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		g.ViewOptions.CameraPos.Y += scrollSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		g.ViewOptions.CameraPos.Y -= scrollSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		g.ViewOptions.CameraPos.X += scrollSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		g.ViewOptions.CameraPos.X -= scrollSpeed
	}
	w, h := ebiten.WindowSize()
	_, dy := ebiten.Wheel()
	old := g.ViewOptions.Scale
	g.ViewOptions.Scale += dy
	g.ViewOptions.Scale = math.Max(g.ViewOptions.Scale, 1)
	g.ViewOptions.Scale = math.Min(g.ViewOptions.Scale, 16)
	ds := g.ViewOptions.Scale / old
	g.ViewOptions.CameraPos.X *= ds
	g.ViewOptions.CameraPos.Y *= ds
	g.ViewOptions.CameraPos.X += float64(w) * (1 - ds) / 2
	g.ViewOptions.CameraPos.Y += float64(h) * (1 - ds) / 2
}
