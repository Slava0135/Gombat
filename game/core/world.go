package core

type World struct {
	Width, Height int
	Floors        [][]*Floor
	Blocks        [][]*Block
}

type Floor struct {
	Name     string
	Passable bool
}

type Block struct {
	Name         string
	Solid        bool
	Destructible bool
}

var (
	Floors = []*Floor{
		{"concrete", true},
		{"grass", true},
		{"road", true},
		{"water", false},
	}
	Blocks = []*Block{
		{"air", false, false},
		{"bricks", true, false},
		{"glass", true, true},
		{"planks", true, true},
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
	return w
}
