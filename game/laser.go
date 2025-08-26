package game

import (
	"space_shooter/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

// Laser representa o objeto do laser dentro do jogo
type Laser struct {
	Image    *ebiten.Image
	Position Vector
}

// NewLaser é responsável por criar uma instância de Laser
func NewLaser(position Vector) *Laser {
	image := assets.LaserSprite

	bounds := image.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfY := float64(bounds.Dy()) / 2

	position.X -= halfW
	position.Y -= halfY

	return &Laser{
		Image:    image,
		Position: position,
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
