package core

type World struct {
	Width, Height int
	Floors        [][]*Floor
	Blocks        [][]*Block
	Stains        [][]bool
}

type Floor struct {
	Name     string
	Passable bool
}

type Block struct {
	Name         string
	Solid        bool
	Destructible bool
	Fragile      bool
}

var (
	Floors = []*Floor{
		{"stones", true},
		{"grass", true},
		{"road", true},
		{"water", false},
		{"sand", true},
	}
	Blocks = []*Block{
		{"air", false, false, false},
		{"bricks", true, false, false},
		{"glass", true, true, true},
		{"planks", true, true, false},
	}
)

func NewEmptyWorld(width, height int) *World {
	w := new(World)
	w.Width = width
	w.Height = height
	w.Blocks = make([][]*Block, width)
	for i := range w.Blocks {
		w.Blocks[i] = make([]*Block, height)
	}
	w.Floors = make([][]*Floor, width)
	for i := range w.Floors {
		w.Floors[i] = make([]*Floor, height)
	}
	w.Stains = make([][]bool, width)
	for i := range w.Stains {
		w.Stains[i] = make([]bool, height)
	}
	return w
}
