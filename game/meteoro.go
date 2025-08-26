package game

import (
	"math/rand"
	"space_shooter/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

// Meteor representa o objeto do meteor dentro do jogo
type Meteor struct {
	Image    *ebiten.Image
	Speed    float64
	Position Vector
}

// NewMeteor é responsável por criar uma instância de Meteor
func NewMeteor() *Meteor {
	image := assets.MeteorSprites[rand.Intn(len(assets.MeteorSprites))]
	speed := rand.Float64() * 13

	position := Vector{
		X: rand.Float64() * screenWidth,
		Y: -100,
	}

	return &Meteor{
		Image:    image,
		Speed:    speed,
		Position: position,
	}
}

// Update é responsável por atualizar a lógica do Meteor
func (meteor *Meteor) Update() {
	meteor.Position.Y += meteor.Speed
}

// Draw é responsável por desenhar o Meteor na tela
func (meteor *Meteor) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}

	options.GeoM.Translate(meteor.Position.X, meteor.Position.Y)

	screen.DrawImage(meteor.Image, options)
}

// Collider determina as dimensões do retângulo de hitbox do meteoro
func (meteor *Meteor) Collider() Rect {
	bounds := meteor.Image.Bounds()

	return NewRect(
		meteor.Position.X,
		meteor.Position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}
