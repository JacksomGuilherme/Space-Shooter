package game

import (
	"math"
	"math/rand"
	"space_shooter/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

// PowerUp representa o objeto do PowerUp dentro do jogo
type PowerUp struct {
	Image    *ebiten.Image
	Position Vector
	Sound    []byte
	Action   func()
}

// NewPowerUp é responsável por criar uma instância de PowerUp
func (game *Game) NewPowerUp(position Vector) *PowerUp {

	powerUp := game.SortPowerUp()

	bounds := powerUp.Image.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfY := float64(bounds.Dy()) / 2

	position.X -= halfW
	position.Y -= halfY

	powerUpSound := assets.ItemPickupSFX

	powerUp.Sound = powerUpSound
	powerUp.Position = position

	return powerUp

}

// Update é responsável por atualizar a lógica do PowerUp
func (powerUp *PowerUp) Update() {
	speed := 3.0

	powerUp.Position.Y += speed
}

// Draw é responsável por desenhar o PowerUp na tela
func (powerUp *PowerUp) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}

	options.GeoM.Translate(powerUp.Position.X, powerUp.Position.Y)

	screen.DrawImage(powerUp.Image, options)
}

// Collider determina as dimensões do retângulo de hitbox do PowerUp
func (powerUp *PowerUp) Collider() Rect {
	bounds := powerUp.Image.Bounds()

	return NewRect(
		powerUp.Position.X,
		powerUp.Position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}

func (game *Game) SortPowerUp() *PowerUp {
	image := assets.HealingSprite
	action := func() {}
	powerUpNumber := rand.Intn(2) + 1
	switch {
	case powerUpNumber == 1:
		action = func() {
			healAmount := 20
			shipHealth := game.Player.Ship.Health
			game.Player.Ship.Health = int(math.Min(float64(shipHealth+healAmount), 100))
		}
	case powerUpNumber == 2:
		image = assets.BlueShieldSprite
		action = func() {
			game.Player.Ship.ShieldActivated = true
		}
	}
	return &PowerUp{
		Image:  image,
		Action: action,
	}
}
