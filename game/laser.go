package game

import (
	"math/rand"
	"space_shooter/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

// Laser representa o objeto do laser dentro do jogo
type Laser struct {
	Image    *ebiten.Image
	Position Vector
	Sound    []byte
}

// NewLaser é responsável por criar uma instância de Laser
func NewLaser(position Vector) *Laser {
	image := assets.LaserSprite

	bounds := image.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfY := float64(bounds.Dy()) / 2

	position.X -= halfW
	position.Y -= halfY

	laserSound := assets.LaserSFX[rand.Intn(len(assets.LaserSFX))]

	return &Laser{
		Image:    image,
		Position: position,
		Sound:    laserSound,
	}
}

// Update é responsável por atualizar a lógica do Laser
func (laser *Laser) Update() {
	speed := 7.0

	laser.Position.Y += -speed
}

// Draw é responsável por desenhar o Laser na tela
func (laser *Laser) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}

	options.GeoM.Translate(laser.Position.X, laser.Position.Y)

	screen.DrawImage(laser.Image, options)
}

// Collider determina as dimensões do retângulo de hitbox do laser
func (laser *Laser) Collider() Rect {
	bounds := laser.Image.Bounds()

	return NewRect(
		laser.Position.X,
		laser.Position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}
