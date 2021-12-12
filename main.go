package main

import (
	"Gombat/core"
	"Gombat/core/assets"
	"github.com/hajimehoshi/ebiten/v2"
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
	if err := ebiten.RunGame(core.NewGame(images)); err != nil {
		log.Fatal(err)
	}
}
