package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gombat/game"
	"log"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Gombat")
	ebiten.SetWindowResizable(true)
	if err := ebiten.RunGame(game.NewGame()); err != nil {
		log.Fatal(err)
	}
}
