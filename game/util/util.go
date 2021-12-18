package util

import (
	"image/color"
	"math"
)

type Position struct {
	X, Y float64
}

var (
	Red    = color.RGBA{0xf0, 0x0, 0x0, 0xff}
	Green  = color.RGBA{0x0, 0xf0, 0x0, 0xff}
	Blue   = color.RGBA{0x0, 0x0, 0xf0, 0xff}
	Cyan   = color.RGBA{0x0, 0xdd, 0xe4, 0xff}
	Yellow = color.RGBA{0xf0, 0xf0, 0x0, 0xff}
	Orange = color.RGBA{0xff, 0xa0, 0x0, 0xff}
)

var TeamColors = []color.RGBA{
	Cyan,
	Orange,
	Red,
	Green,
}

func (p *Position) DistanceTo(other Position) float64 {
	return math.Sqrt(math.Pow(p.X-other.X, 2) + math.Pow(p.Y-other.Y, 2))
}
