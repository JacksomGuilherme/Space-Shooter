package main

import (
	"log"
	"space_shooter/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	g := game.NewGame()
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Space Shooter")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
