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
		{Name: "stones", Passable: true},
		{Name: "grass", Passable: true},
		{Name: "road", Passable: true},
		{Name: "water", Passable: false},
		{Name: "sand", Passable: true},
	}
	Blocks = []*Block{
		{Name: "air", Solid: false, Destructible: false, Fragile: false},
		{Name: "bricks", Solid: true, Destructible: false, Fragile: false},
		{Name: "glass", Solid: true, Destructible: true, Fragile: true},
		{Name: "planks", Solid: true, Destructible: true, Fragile: false},
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
