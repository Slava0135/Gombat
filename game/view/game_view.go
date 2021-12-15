package view

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gombat/game/core"
	"image/color"
)

func DrawGameState(img *ebiten.Image, state *core.GameState) {
	img.Fill(color.White)
	drawWorld(img, state.World)
}
