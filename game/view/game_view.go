package view

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gombat/game/assets"
	"gombat/game/core"
	"image/color"
)

const gopSize = tileSize * tileSize * core.GopSize / 11

func DrawGameState(img *ebiten.Image, state *core.GameState) {
	img.Fill(color.White)
	drawWorld(img, state.World)
	drawObjects(img, state)
}

func drawObjects(img *ebiten.Image, state *core.GameState) {
	drawGops(img, state)
}

func drawGops(img *ebiten.Image, state *core.GameState) {
	op := new(ebiten.DrawImageOptions)
	op.GeoM.Scale(gopSize, gopSize)
	op.GeoM.Translate(4*drawScale*tileSize, 6*drawScale*tileSize)
	for _, team := range state.Teams {
		gopImg := assets.Images["gop"]
		for range team.Gops {
			img.DrawImage(gopImg, op)
		}
	}
}
