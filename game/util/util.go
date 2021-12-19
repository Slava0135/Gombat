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

func RayTrace(grid [][]bool, from, to Position) (x, y int, ok bool) {

	dx, dy := math.Abs(from.X-to.X), math.Abs(from.Y-to.Y)
	x, y = int(math.Floor(from.X)), int(math.Floor(from.Y))

	var xInc, yInc int
	var err float64

	if dx == 0 {
		xInc = 0
		err = math.Inf(1)
	} else if to.X > from.X {
		xInc = 1
		err = (math.Floor(from.X) + 1 - from.X) * dy
	} else {
		xInc = -1
		err = (from.X - math.Floor(from.X)) * dy
	}

	if dy == 0 {
		yInc = 0
		err -= math.Inf(1)
	} else if to.Y > from.Y {
		yInc = 1
		err -= (math.Floor(from.Y) + 1 - from.Y) * dx
	} else {
		yInc = -1
		err -= (from.Y - math.Floor(from.Y)) * dx
	}

	xEnd, yEnd := int(math.Floor(to.X)), int(math.Floor(to.Y))
	for x != xEnd || y != yEnd {
		if grid[x][y] {
			return x, y, true
		}
		if err > 0 {
			y += yInc
			err -= dx
		} else {
			x += xInc
			err += dy
		}
	}

	return x, y, grid[x][y]
}
