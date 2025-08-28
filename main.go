package main

import (
	"log"
	"space_shooter/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	g := game.NewGame()
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Space Shooter")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
