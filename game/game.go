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
	viewport         viewport
}

var (
	BackgroundImage *ebiten.Image
)

func init() {
	BackgroundImage = assets.BackgroundSprite
}

type viewport struct {
	x16 int
	y16 int
}

// Move é responsável por mover a viewport com a imagem do fundo
func (p *viewport) Move() {
	s := BackgroundImage.Bounds().Size()
	p.y16 += (s.Y / 512) * -1

	if p.y16 <= 0 {
		p.y16 = s.Y * 16
	}
}

// Position determina a posição do viewport
func (p *viewport) Position() (int, int) {
	return p.x16, p.y16
}

func NewViewport() viewport {
	s := BackgroundImage.Bounds().Size()
	return viewport{
		x16: 0,
		y16: s.Y * 16, // começa no "pé" da imagem
	}
}

func NewGame() *Game {
	g := &Game{
		MeteorSpawnTimer: NewTimer(24),
		viewport:         NewViewport(),
	}
	player := NewPlayer(g)
	g.Player = player
	return g
}

// Update é responsável por atualizar a logica do jogo
func (game *Game) Update() error {
	game.Player.Update()

	game.viewport.Move()

	for _, laser := range game.Lasers {
		laser.Update()
	}

	game.MeteorSpawnTimer.Update()
	if game.MeteorSpawnTimer.IsReady() {
		game.MeteorSpawnTimer.Reset()

		meteor := NewMeteor()

		game.Meteors = append(game.Meteors, meteor)
	}

	for _, meteor := range game.Meteors {
		meteor.Update()
	}

	for _, meteor := range game.Meteors {
		if meteor.Collider().Intersects(game.Player.Collider()) {
			assets.PlaySFX(game.Player.Sound)
			game.Reset()
		}
	}

	for i, meteor := range game.Meteors {
		for j, laser := range game.Lasers {
			if meteor.Collider().Intersects(laser.Collider()) {
				assets.PlaySFX(meteor.Sound)
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
	_, y16 := game.viewport.Position()
	offsetY := float64(-y16) / 16

	_, h := BackgroundImage.Bounds().Dx(), BackgroundImage.Bounds().Dy()

	for j := 0; j < 2; j++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(0, float64(h*j)+offsetY)
		screen.DrawImage(BackgroundImage, op)
	}
	for _, laser := range game.Lasers {
		laser.Draw(screen)
	}

	game.Player.Draw(screen)

	for _, meteor := range game.Meteors {
		meteor.Draw(screen)
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
func (game *Game) Reset() {

	game.Player = NewPlayer(game)
	game.Meteors = nil
	game.Lasers = nil
	game.MeteorSpawnTimer.Reset()
	game.Score = 0
}
