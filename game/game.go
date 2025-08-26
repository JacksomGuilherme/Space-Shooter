package game

import (
	"fmt"
	"image/color"
	"space_shooter/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

// Game representa o objeto do jogo
type Game struct {
	Player           *Player
	Lasers           []*Laser
	Meteors          []*Meteor
	MeteorSpawnTimer *Timer
	Score            int
}

func NewGame() *Game {
	g := &Game{
		MeteorSpawnTimer: NewTimer(24),
	}
	player := NewPlayer(g)
	g.Player = player
	return g
}

// Update é responsável por atualizar a logica do jogo
func (game *Game) Update() error {
	game.Player.Update()

	for _, laser := range game.Lasers {
		laser.Update()
	}

	game.MeteorSpawnTimer.Update()
	if game.MeteorSpawnTimer.IsReady() {
		game.MeteorSpawnTimer.Reset()

		meteor := NewMeteor()

		game.Meteors = append(game.Meteors, meteor)
	}

	for _, meteoro := range game.Meteors {
		meteoro.Update()
	}

	for _, meteoro := range game.Meteors {
		if meteoro.Collider().Intersects(game.Player.Collider()) {
			game.Rest()
		}
	}

	for i, meteoro := range game.Meteors {
		for j, laser := range game.Lasers {
			if meteoro.Collider().Intersects(laser.Collider()) {
				game.Meteors = append(game.Meteors[:i], game.Meteors[i+1:]...)
				game.Lasers = append(game.Lasers[:j], game.Lasers[j+1:]...)

				game.Score++
				break
			}
		}
	}

	return nil
}

// Draw é responsável por desenhar os objetos na tela
func (game *Game) Draw(screen *ebiten.Image) {
	game.Player.Draw(screen)

	for _, laser := range game.Lasers {
		laser.Draw(screen)
	}

	for _, meteoro := range game.Meteors {
		meteoro.Draw(screen)
	}

	text.Draw(screen, fmt.Sprintf("Pontos: %d", game.Score), assets.FontUi, 20, 100, color.White)
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// AddLasers adiciona um novo laser ao slice de lasers
func (game *Game) AddLasers(laser *Laser) {
	game.Lasers = append(game.Lasers, laser)
}

// Rest reinicia o jogo do zero
func (game *Game) Rest() {

	game.Player = NewPlayer(game)
	game.Meteors = nil
	game.Lasers = nil
	game.MeteorSpawnTimer.Reset()
	game.Score = 0
}
