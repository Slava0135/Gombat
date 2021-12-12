package assets

import (
	"bytes"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"log"
)

type Images struct {
	Bricks   *ebiten.Image
	Concrete *ebiten.Image
	Glass    *ebiten.Image
	Gop      *ebiten.Image
	Grass    *ebiten.Image
	Planks   *ebiten.Image
	Road     *ebiten.Image
	Water    *ebiten.Image
}

func LoadImages() *Images {
	images := new(Images)
	images.Bricks = loadImage(bricks)
	images.Concrete = loadImage(concrete)
	images.Glass = loadImage(glass)
	images.Gop = loadImage(gop)
	images.Grass = loadImage(grass)
	images.Planks = loadImage(planks)
	images.Road = loadImage(road)
	images.Water = loadImage(water)
	return images
}

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

func loadImage(b []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}
	return ebiten.NewImageFromImage(img)
}
