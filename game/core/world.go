package core

type World struct {
	Height, Width int
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
		{"bricks", true, false},
		{"glass", true, true},
		{"planks", true, true},
	}
)

func NewEmptyWorld(height, width int) *World {
	w := new(World)
	w.Height = height
	w.Width = width
	w.Blocks = make([][]*Block, height)
	for i := range w.Blocks {
		w.Blocks[i] = make([]*Block, width)
	}
	w.Floors = make([][]*Floor, height)
	for i := range w.Floors {
		w.Floors[i] = make([]*Floor, width)
	}
	return w
}
