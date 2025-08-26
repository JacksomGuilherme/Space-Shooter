package game

import (
	"space_shooter/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

// Player representa o objeto do jogador dentro do jogo
type Player struct {
	Image             *ebiten.Image
	Position          Vector
	Game              *Game
	LaserLoadingTimer *Timer
}

// NewPlayer é responsável por criar uma instância de Player
func NewPlayer(game *Game) *Player {
	image := assets.PlayerSprite

	bounds := image.Bounds()
	halfW := float64(bounds.Dx()) / 2

	position := Vector{
		X: screenWidth/2 - halfW,
		Y: 500,
	}

	return &Player{
		Image:             image,
		Position:          position,
		Game:              game,
		LaserLoadingTimer: NewTimer(12),
	}
}

// Update é responsável por atualizar a lógica do Player
func (player *Player) Update() {

	speed := 6.0

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		player.Position.X -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		player.Position.X += speed
	}

	player.LaserLoadingTimer.Update()
	if ebiten.IsKeyPressed(ebiten.KeySpace) && player.LaserLoadingTimer.IsReady() {
		player.LaserLoadingTimer.Reset()

		bounds := player.Image.Bounds()
		halfW := float64(bounds.Dx()) / 2
		halfY := float64(bounds.Dy()) / 2

		spawnPosition := Vector{
			player.Position.X + halfW,
			player.Position.Y - halfY/2,
		}
		laser := NewLaser(spawnPosition)
		player.Game.AddLasers(laser)
	}
}

// Drawte é responsável por desenhar Player na tela
func (player *Player) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}

	options.GeoM.Translate(player.Position.X, player.Position.Y)

	screen.DrawImage(player.Image, options)
}

// Collider determina as dimensões do retângulo de hitbox do player
func (player *Player) Collider() Rect {
	bounds := player.Image.Bounds()

	return NewRect(
		player.Position.X,
		player.Position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}
