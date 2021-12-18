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
	//go:embed images/floors/stones.png
	stones []byte
	//go:embed images/blocks/glass.png
	glass []byte
	//go:embed images/gop.png
	gop []byte
	//go:embed images/gop_team.png
	gopTeam []byte
	//go:embed images/floors/grass.png
	grass []byte
	//go:embed images/blocks/planks.png
	planks []byte
	//go:embed images/floors/road.png
	road []byte
	//go:embed images/floors/water.png
	water []byte
)

var Images = map[string]*ebiten.Image{
	"bricks":  loadImage(bricks),
	"stones":  loadImage(stones),
	"glass":   loadImage(glass),
	"gop":     loadImage(gop),
	"gopteam": loadImage(gopTeam),
	"grass":   loadImage(grass),
	"planks":  loadImage(planks),
	"road":    loadImage(road),
	"water":   loadImage(water),
}

func loadImage(b []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}
	return ebiten.NewImageFromImage(img)
}
