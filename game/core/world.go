package core

type World struct {
	Height, Width int
	Floors        [][]*Floor
	Blocks        [][]*Block
}

type Floor struct {
	Id       byte
	Name     string
	Passable bool
}

type Block struct {
	Id           byte
	Name         string
	Solid        bool
	Destructible bool
}

var (
	Concrete = &Floor{0, "concrete", true}
	Grass    = &Floor{1, "grass", true}
	Road     = &Floor{2, "road", true}
	Water    = &Floor{3, "water", false}
)

var (
	Bricks = &Block{0, "bricks", true, false}
	Glass  = &Block{1, "glass", true, true}
	Planks = &Block{2, "planks", true, true}
)

var (
	Floors = []*Floor{Concrete, Grass, Road, Water}
	Blocks = []*Block{Bricks, Glass, Planks}
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
