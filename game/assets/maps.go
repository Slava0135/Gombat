package assets

import (
	"errors"
	"gombat/game/core"
	"gombat/game/util"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func LoadMap(name string) (world *core.World, t []*core.Team, err error) {

	defer func() {
		if r := recover(); r != nil {
			world, err = nil, errors.New("wrong map format")
		}
	}()

	path := filepath.Join("game", "assets", "maps", name)
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, errors.New("wrong map format")
	}

	chunks := strings.Split(string(bytes), "\n\n")

	floors := strings.Split(chunks[0], "\n")
	blocks := strings.Split(chunks[1], "\n")
	teams := strings.Split(chunks[2], "\n")

	width := len(floors[0])
	height := len(floors)

	world = core.NewEmptyWorld(width, height)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			world.Floors[x][y] = core.Floors[floors[y][x]-'0']
			world.Blocks[x][y] = core.Blocks[blocks[y][x]-'0']
		}
	}

	t = make([]*core.Team, 0, len(teams))
	for i := range teams {
		team := &core.Team{Id: i, Gops: make([]*core.Gop, 0, 4), DeadGops: make([]util.Vec2, 0, 4)}
		coords := strings.Split(teams[i], " ")
		for j := range coords {
			rawCoord := strings.Split(coords[j], "-")
			x, _ := strconv.ParseFloat(rawCoord[0], 64)
			y, _ := strconv.ParseFloat(rawCoord[1], 64)
			gop := &core.Gop{Health: core.GopHealth, Team: team, Pos: util.Vec2{X: x, Y: y}}
			team.Gops = append(team.Gops, gop)
		}
		t = append(t, team)
	}

	return world, t, nil
}
