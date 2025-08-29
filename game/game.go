package game

import (
	"space_shooter/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game representa o objeto do jogo
type Game struct {
	mode              Mode
	Player            *Player
	Lasers            []*Laser
	Meteors           []*Meteor
	PowerUps          []*PowerUp
	Menu              *Menu
	MeteorSpawnTimer  *Timer
	PowerUpSpawnTimer *Timer
	Score             int
	MaxScore          int
	viewport          viewport
	gameOverCount     int
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
	ModePause
	ModeShipSelection
)

type viewport struct {
	x16    int
	y16    int
	moving bool
}

// Move é responsável por mover a viewport com a imagem do fundo
func (p *viewport) Move() {
	if !p.moving {
		return
	}
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
		x16:    0,
		y16:    s.Y * 16, // começa no "pé" da imagem
		moving: false,
	}
}

func NewGame() *Game {
	g := &Game{
		MeteorSpawnTimer:  NewTimer(24),
		PowerUpSpawnTimer: NewTimer(150),
		viewport:          NewViewport(),
	}

	player := NewPlayer(g)
	g.Player = player

	g.NewTitleMode()

	return g
}

// Update é responsável por atualizar a logica do jogo
func (game *Game) Update() error {
	switch game.mode {
	case ModeTitle:
		game.UpdateTitleMode()
	case ModeGame:
		game.viewport.moving = true
		game.UpdateGameMode()
	case ModeGameOver:
		game.UpdateGameOverMode()
	case ModePause:
		game.UpdatePauseMode()
	case ModeShipSelection:
		game.UpdateShipSelectionMode()
	}

	return nil
}

// Draw é responsável por desenhar os objetos na tela
func (game *Game) Draw(screen *ebiten.Image) {
	game.viewport.Move()

	_, y16 := game.viewport.Position()
	offsetY := float64(-y16) / 16

	_, h := BackgroundImage.Bounds().Dx(), BackgroundImage.Bounds().Dy()

	for j := 0; j < 2; j++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(0, float64(h*j)+offsetY)
		screen.DrawImage(BackgroundImage, op)
	}

	switch game.mode {
	case ModeTitle:
		game.DrawTitleMode(screen)
	case ModeGame:
		game.DrawGameMode(screen)
	case ModeGameOver:
		game.DrawGameOverMode(screen)
	case ModePause:
		game.DrawGameMode(screen)
		game.DrawPauseMode(screen)
	case ModeShipSelection:
		game.DrawShipSelectionMode(screen)
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
	if game.MaxScore < game.Score {
		game.MaxScore = game.Score
	}
	game.viewport = NewViewport()
	game.Player = NewPlayer(game)
	game.Meteors = nil
	game.Lasers = nil
	game.MeteorSpawnTimer.Reset()
	game.Menu.SelectedItemIndex = 0
	game.Score = 0
}
