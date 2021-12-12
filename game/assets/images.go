package assets

import (
	"bytes"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"log"
)

var (
	//go:embed images/blocks/bricks.png
	bricks []byte
	//go:embed images/floors/concrete.png
	concrete []byte
	//go:embed images/blocks/glass.png
	glass []byte
	//go:embed images/gop.png
	gop []byte
	//go:embed images/floors/grass.png
	grass []byte
	//go:embed images/blocks/planks.png
	planks []byte
	//go:embed images/floors/road.png
	road []byte
	//go:embed images/floors/water.png
	water []byte
)

func LoadImages() *map[string]*ebiten.Image {
	return &map[string]*ebiten.Image{
		"bricks":   loadImage(bricks),
		"concrete": loadImage(concrete),
		"glass":    loadImage(glass),
		"gop":      loadImage(gop),
		"grass":    loadImage(grass),
		"planks":   loadImage(planks),
		"road":     loadImage(road),
		"water":    loadImage(water),
	}
}

func loadImage(b []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}
	return ebiten.NewImageFromImage(img)
}
