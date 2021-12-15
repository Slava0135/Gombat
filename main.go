package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gombat/game"
	"log"
)

const (
	screenWidth  = 1280
	screenHeight = 960
)

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Gombat")
	ebiten.SetWindowResizable(true)
	if err := ebiten.RunGame(game.NewGame()); err != nil {
		log.Fatal(err)
	}
}
