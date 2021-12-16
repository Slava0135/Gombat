package view

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gombat/game/assets"
	"gombat/game/core"
	"gombat/game/util"
	"image/color"
)

type ViewOptions struct {
	CameraPos *util.Position
	Scale     float64
}

const gopSize = tileSize * core.GopSize / 11

func DrawGameState(img *ebiten.Image, viewOptions *ViewOptions, state *core.GameState) {
	img.Fill(color.White)
	drawWorld(img, viewOptions, state.World)
	drawObjects(img, viewOptions, state)
}

func drawObjects(img *ebiten.Image, viewOptions *ViewOptions, state *core.GameState) {
	drawGops(img, viewOptions, state)
}

func drawGops(img *ebiten.Image, viewOptions *ViewOptions, state *core.GameState) {
	op := new(ebiten.DrawImageOptions)
	op.GeoM.Scale(viewOptions.Scale, viewOptions.Scale)
	op.GeoM.Scale(gopSize, gopSize)
	op.GeoM.Translate(viewOptions.CameraPos.X, viewOptions.CameraPos.Y)
	op.GeoM.Translate(4*viewOptions.Scale*tileSize, 6*viewOptions.Scale*tileSize)
	for _, team := range state.Teams {
		gopImg := assets.Images["gop"]
		for range team.Gops {
			img.DrawImage(gopImg, op)
		}
	}
}
