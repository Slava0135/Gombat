package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gombat/game/assets"
	"gombat/game/core"
	"gombat/game/util"
	"gombat/game/view"
	"math"
)

const scrollSpeed = 32

type Game struct {
	GameState     *core.GameState
	ViewOptions   *view.ViewOptions
	SelectOptions *SelectOptions
}

type SelectOptions struct {
	GopSelected *core.Gop
}

func NewGame() *Game {
	g := new(Game)
	g.GameState = core.NewGameState(2, 4)
	m, _ := assets.LoadMap("river")
	g.GameState.World = m
	g.ViewOptions = &view.ViewOptions{
		CameraPos: &util.Vec2{},
		Scale:     4,
	}
	g.SelectOptions = new(SelectOptions)
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
	g.updateCameraPosition()
	g.updateScale()
	g.updateSelection()
}

func (g *Game) updateCameraPosition() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		g.ViewOptions.CameraPos.Y -= scrollSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		g.ViewOptions.CameraPos.Y += scrollSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		g.ViewOptions.CameraPos.X -= scrollSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		g.ViewOptions.CameraPos.X += scrollSpeed
	}
}

func (g *Game) updateScale() {
	w, h := ebiten.WindowSize()
	_, dy := ebiten.Wheel()
	oldScale := g.ViewOptions.Scale
	g.ViewOptions.Scale += dy
	g.ViewOptions.Scale = math.Max(g.ViewOptions.Scale, 1)
	g.ViewOptions.Scale = math.Min(g.ViewOptions.Scale, 16)
	ds := g.ViewOptions.Scale / oldScale
	g.ViewOptions.CameraPos.X *= ds
	g.ViewOptions.CameraPos.Y *= ds
	g.ViewOptions.CameraPos.X += float64(w) * (ds - 1) / 2
	g.ViewOptions.CameraPos.Y += float64(h) * (ds - 1) / 2
}

func (g *Game) updateSelection() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		worldX := float64(x) + g.ViewOptions.CameraPos.X
		worldY := float64(y) + g.ViewOptions.CameraPos.Y
		worldX /= view.TileImgSize * g.ViewOptions.Scale
		worldY /= view.TileImgSize * g.ViewOptions.Scale
		pos := util.Vec2{worldX, worldY}
		if gop := g.SelectOptions.GopSelected; gop == nil {
			g.SelectOptions.GopSelected = g.GameState.SelectGop(pos)
		} else {
			if ok := g.GameState.World.CanMoveGop(gop.Pos, pos); ok {
				gop.MoveGop(pos)
			}
		}
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		g.SelectOptions.GopSelected = nil
	}
}
