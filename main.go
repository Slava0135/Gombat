package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gombat/game"
	"gombat/game/assets"
	"log"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

func main() {
	images := assets.LoadImages()
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Gombat")
	ebiten.SetWindowResizable(true)
	if err := ebiten.RunGame(game.NewGame(images)); err != nil {
		log.Fatal(err)
	}
}
