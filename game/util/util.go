package util

import (
	"image/color"
	"math"
)

type Vec2 struct {
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

func (v *Vec2) DistanceTo(other Vec2) float64 {
	return math.Sqrt(math.Pow(v.X-other.X, 2) + math.Pow(v.Y-other.Y, 2))
}

func (v *Vec2) IsInSquareBounds(center Vec2, dist float64) bool {
	return math.Abs(v.X-center.X) < dist && math.Abs(v.Y-center.Y) < dist
}

func (v *Vec2) Add(other Vec2) Vec2 {
	return Vec2{v.X + other.X, v.Y + other.Y}
}

func (v *Vec2) Rotate(a float64) Vec2 {
	return Vec2{
		math.Cos(a)*v.X - math.Sin(a)*v.Y,
		math.Sin(a)*v.X + math.Cos(a)*v.Y,
	}
}

func (v *Vec2) Norm() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}

func RayTrace(grid [][]bool, from, to Vec2) (x, y int, collided bool) {

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

	w, h := len(grid), len(grid[0])
	xEnd, yEnd := int(math.Floor(to.X)), int(math.Floor(to.Y))
	for x != xEnd || y != yEnd {
		if x >= 0 && x < w && y >= 0 && y < h && grid[x][y] {
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

	return x, y, x >= 0 && x < w && y >= 0 && y < h && grid[x][y]
}
