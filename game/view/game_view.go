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
	drawWorld(img, viewOptions, state.World)
	drawObjects(img, viewOptions, state)
}

func drawObjects(img *ebiten.Image, viewOptions *ViewOptions, state *core.GameState) {
	drawGops(img, viewOptions, state)
}

func drawGops(img *ebiten.Image, viewOptions *ViewOptions, state *core.GameState) {

	gopImg := assets.Images["gop"]
	gopTeamImg := assets.Images["gopteam"]

	commonOp := new(ebiten.DrawImageOptions)
	commonOp.GeoM.Scale(viewOptions.Scale, viewOptions.Scale)
	commonOp.GeoM.Scale(gopSize, gopSize)
	commonOp.GeoM.Translate(viewOptions.CameraPos.X, viewOptions.CameraPos.Y)
	op := new(ebiten.DrawImageOptions)
	for _, team := range state.Teams {
		colorOp := new(ebiten.DrawImageOptions)

		c := teamColors[team.Id]
		colorOp.ColorM.Scale(0, 0, 0, 1)
		r := float64(c.R) / 0xff
		g := float64(c.G) / 0xff
		b := float64(c.B) / 0xff
		colorOp.ColorM.Translate(r, g, b, 0)

		gopImgColored := ebiten.NewImageFromImage(gopImg)
		gopImgColored.DrawImage(gopTeamImg, colorOp)
		for _, gop := range team.Gops {
			op.GeoM.Reset()
			op.GeoM.Concat(commonOp.GeoM)
			op.GeoM.Translate(gop.Pos.X*viewOptions.Scale*tileSize, gop.Pos.Y*viewOptions.Scale*tileSize)
			img.DrawImage(gopImgColored, op)
		}
	}
}

var teamColors = []color.RGBA{
	util.Cyan,
	util.Red,
	util.Green,
	util.Blue,
}
