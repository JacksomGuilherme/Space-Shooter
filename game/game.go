package game

import (
	"fmt"
	"image/color"
	"space_shooter/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

// Game representa o objeto do jogo
type Game struct {
	mode             Mode
	Player           *Player
	Lasers           []*Laser
	Meteors          []*Meteor
	MeteorSpawnTimer *Timer
	Score            int
	viewport         viewport
	gameOverCount    int
}

var (
	BackgroundImage *ebiten.Image
)

func init() {
	BackgroundImage = assets.BackgroundSprite
}

type Mode int

const (
	ModeTitle = iota
	ModeGame
	ModeGameOver
)

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
	menuSound := assets.MenuSFX
	switch game.mode {
	case ModeTitle:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			assets.PlaySFX(menuSound, 1)
			game.mode = ModeGame
		}
	case ModeGame:
		game.UpdateGameMode()
	case ModeGameOver:
		if game.gameOverCount > 0 {
			game.gameOverCount--
		}
		if game.gameOverCount == 0 && inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			game.Reset()
			assets.PlaySFX(menuSound, 1)
			game.mode = ModeTitle
		}
	}

	return nil
}

// Draw é responsável por desenhar os objetos na tela
func (game *Game) Draw(screen *ebiten.Image) {
	game.viewport.Move()

	switch game.mode {
	case ModeTitle:
		game.DrawTitleMode(screen)
	case ModeGame:
		game.DrawGameMode(screen)
	case ModeGameOver:
		game.DrawGameOverMode(screen)
	}
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

func (game *Game) UpdateGameMode() {
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

	for _, meteor := range game.Meteors {
		meteor.Update()
	}

	if game.Player.Health <= 0 {
		assets.PlaySFX(game.Player.DeathSound, 1)
		game.mode = ModeGameOver
		game.gameOverCount = 30
		game.Reset()
	}

	for i, meteor := range game.Meteors {
		if meteor.Collider().Intersects(game.Player.Collider()) {
			if !meteor.Hit {
				assets.PlaySFX(game.Player.HitSound, 1)
				game.Player.Health -= meteor.DamageByClass()
				meteor.Hit = true
				game.Meteors = append(game.Meteors[:i], game.Meteors[i+1:]...)
				break
			}
		}
	}

	for i, meteor := range game.Meteors {
		for j, laser := range game.Lasers {
			if meteor.Collider().Intersects(laser.Collider()) {
				assets.PlaySFX(meteor.Sound, 1)
				game.Meteors = append(game.Meteors[:i], game.Meteors[i+1:]...)
				game.Lasers = append(game.Lasers[:j], game.Lasers[j+1:]...)
				game.Score++
				break
			}
		}
	}
}

func (game *Game) DrawGameMode(screen *ebiten.Image) {
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

	text.Draw(screen, fmt.Sprintf("Points: %d", game.Score), assets.FontUi, 20, 100, color.White)
}

func (game *Game) DrawTitleMode(screen *ebiten.Image) {
	_, y16 := game.viewport.Position()
	offsetY := float64(-y16) / 16

	_, h := BackgroundImage.Bounds().Dx(), BackgroundImage.Bounds().Dy()

	for j := 0; j < 2; j++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(0, float64(h*j)+offsetY)
		screen.DrawImage(BackgroundImage, op)
	}

	game.Player.Draw(screen)

	titleText := "Space Shooter"
	bounds, _ := font.BoundString(assets.FontUi, titleText)
	textHalfWidth := bounds.Max.X / 2
	text.Draw(screen, titleText, assets.FontUi, (screenWidth/2 - textHalfWidth.Round()), 3*titleFontSize, color.White)

	texts := "PRESS SPACE KEY"
	bounds, _ = font.BoundString(assets.FontUi, texts)
	textHalfWidth = bounds.Max.X / 2
	text.Draw(screen, texts, assets.FontUi, (screenWidth/2 - textHalfWidth.Round()), screenHeight-titleFontSize*5, color.White)

}

func (game *Game) DrawGameOverMode(screen *ebiten.Image) {
	_, y16 := game.viewport.Position()
	offsetY := float64(-y16) / 16

	_, h := BackgroundImage.Bounds().Dx(), BackgroundImage.Bounds().Dy()

	for j := 0; j < 2; j++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(0, float64(h*j)+offsetY)
		screen.DrawImage(BackgroundImage, op)
	}

	titleText := "GAME OVER!"
	bounds, _ := font.BoundString(assets.FontUi, titleText)
	textHalfWidth := bounds.Max.X / 2
	textHalfHeight := bounds.Max.Y / 2
	text.Draw(
		screen,
		titleText,
		assets.FontUi,
		(screenWidth/2 - textHalfWidth.Round()),
		(screenHeight/2 - textHalfHeight.Round()),
		color.White)

}
