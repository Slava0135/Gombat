package assets

import (
	"errors"
	"gombat/game/core"
	"os"
	"path/filepath"
	"strings"
)

func LoadMap(name string) (world *core.World, err error) {

	defer func() {
		if r := recover(); r != nil {
			world, err = nil, errors.New("wrong map format")
		}
	}()

	path := filepath.Join("game", "assets", "maps", name)
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.New("wrong map format")
	}

	chunks := strings.Split(string(bytes), "\n\n")

	floors := strings.Split(chunks[0], "\n")
	blocks := strings.Split(chunks[1], "\n")

	width := len(floors[0])
	height := len(floors)

	world = core.NewEmptyWorld(width, height)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			world.Floors[x][y] = core.Floors[floors[y][x]-'0']
			world.Blocks[x][y] = core.Blocks[blocks[y][x]-'0']
		}
	}
	return world, nil
}
