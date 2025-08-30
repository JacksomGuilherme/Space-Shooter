package game

import (
	"space_shooter/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	laserSpeed  = 7.0
	image       = assets.LaserSprite
	laserSound  = "laser"
	laserBounds = assets.LaserSprite.Bounds()
	laserHalfW  = float64(laserBounds.Dx()) / 2
	laserHalfH  = float64(laserBounds.Dy()) / 2
)

// Laser representa o objeto do laser dentro do jogo
type Laser struct {
	Image    *ebiten.Image
	Position Vector
	Sound    string
}

// NewLaser é responsável por criar uma instância de Laser
func NewLaser(position Vector) *Laser {

	position.X -= laserHalfW
	position.Y -= laserHalfH

	return &Laser{
		Image:    image,
		Position: position,
		Sound:    laserSound,
	}
}

// Update é responsável por atualizar a lógica do Laser
func (laser *Laser) Update() {
	laser.Position.Y += -laserSpeed
}

// Draw é responsável por desenhar o Laser na tela
func (laser *Laser) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}

	options.GeoM.Translate(laser.Position.X, laser.Position.Y)

	screen.DrawImage(laser.Image, options)
}

// Collider determina as dimensões do retângulo de hitbox do laser
func (laser *Laser) Collider() Rect {
	laserBounds := laser.Image.Bounds()

	return NewRect(
		laser.Position.X,
		laser.Position.Y,
		float64(laserBounds.Dx()),
		float64(laserBounds.Dy()),
	)
}
