package view

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gombat/game/assets"
	"gombat/game/core"
)

const drawScale = 4
const tileSize = 8

func drawWorld(img *ebiten.Image, world *core.World) {
	drawFloors(img, world)
	drawBlocks(img, world)
}

func drawFloors(img *ebiten.Image, world *core.World) {
	op := new(ebiten.DrawImageOptions)
	for x := 0; x < world.Width; x++ {
		op.GeoM.Reset()
		op.GeoM.Scale(drawScale, drawScale)
		op.GeoM.Translate(float64(x*drawScale*tileSize), 0)
		for y := 0; y < world.Height; y++ {
			floor := world.Floors[x][y]
			if floorImg, ok := assets.Images[floor.Name]; ok {
				img.DrawImage(floorImg, op)
			}
			op.GeoM.Translate(0, drawScale*tileSize)
		}
	}
}

func drawBlocks(img *ebiten.Image, world *core.World) {
	op := new(ebiten.DrawImageOptions)
	for x := 0; x < world.Width; x++ {
		op.GeoM.Reset()
		op.GeoM.Scale(drawScale, drawScale)
		op.GeoM.Translate(float64(x*drawScale*tileSize), 0)
		for y := 0; y < world.Height; y++ {
			block := world.Blocks[x][y]
			if blockImg, ok := assets.Images[block.Name]; ok {
				img.DrawImage(blockImg, op)
			}
			op.GeoM.Translate(0, drawScale*tileSize)
		}
	}
}
