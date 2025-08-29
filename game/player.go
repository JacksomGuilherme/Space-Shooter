package game

import (
	"space_shooter/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

// Player representa o objeto do jogador dentro do jogo
type Player struct {
	Ship              Ship
	Position          Vector
	Game              *Game
	DeathSound        []byte
	HitSound          []byte
	LaserLoadingTimer *Timer
}

type Ship struct {
	Image           *ebiten.Image
	Health          int
	Laser           *ebiten.Image
	ShieldActivated bool
	ShieldTimer     *Timer
}

// NewPlayer é responsável por criar uma instância de Player
func NewPlayer(game *Game) *Player {
	ship := Ship{
		Image:       assets.PlayerSpriteBlue,
		Health:      100,
		ShieldTimer: NewTimer(600),
	}

	if game.Player != nil && game.Player.Ship.Image != nil {
		ship.Image = game.Player.Ship.Image
	}

	bounds := ship.Image.Bounds()
	halfW := float64(bounds.Dx()) / 2

	position := Vector{
		X: screenWidth/2 - halfW,
		Y: 500,
	}

	deathSound := assets.PlayerDeathSFX
	hitSound := assets.PlayerHitSFX

	return &Player{
		Ship:              ship,
		Position:          position,
		Game:              game,
		DeathSound:        deathSound,
		HitSound:          hitSound,
		LaserLoadingTimer: NewTimer(12),
	}
}

// Update é responsável por atualizar a lógica do Player
func (player *Player) Update() {

	speed := 6.0
	bounds := player.Ship.Image.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfY := float64(bounds.Dy()) / 2

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if (player.Position.X - speed) <= 0-halfW {
			return
		}

		player.Position.X -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if (player.Position.X + speed) >= screenWidth-halfW {
			return
		}
		player.Position.X += speed
	}

	player.LaserLoadingTimer.Update()
	if ebiten.IsKeyPressed(ebiten.KeySpace) && player.LaserLoadingTimer.IsReady() {
		player.LaserLoadingTimer.Reset()

		spawnPosition := Vector{
			player.Position.X + halfW,
			player.Position.Y - halfY/2,
		}
		laser := NewLaser(spawnPosition)
		assets.PlaySFX(laser.Sound, 1)
		player.Game.AddLasers(laser)
	}

	if player.Ship.ShieldActivated {
		player.Ship.ShieldTimer.Update()
		if player.Ship.ShieldTimer.IsReady() {
			player.Ship.ShieldTimer.Reset()

			player.Ship.ShieldActivated = false
		}
	}
}

// Drawte é responsável por desenhar Player na tela
func (player *Player) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}

	options.GeoM.Translate(player.Position.X, player.Position.Y)

	screen.DrawImage(player.Ship.Image, options)

	if player.Ship.ShieldActivated {
		image := assets.ShieldSprite
		bounds := image.Bounds()
		halfW := float64(bounds.Dx()) / 2
		halfH := float64(bounds.Dy()) / 2

		shipBounds := player.Ship.Image.Bounds()
		shipHalfW := float64(shipBounds.Dx()) / 2
		shipHalfH := float64(shipBounds.Dy()) / 2

		options = &ebiten.DrawImageOptions{}
		options.GeoM.Translate(
			player.Position.X+shipHalfW-halfW,
			player.Position.Y+shipHalfH-halfH,
		)
		screen.DrawImage(image, options)
	}

}

// Collider determina as dimensões do retângulo de hitbox do player
func (player *Player) Collider() Rect {
	bounds := player.Ship.Image.Bounds()

	return NewRect(
		player.Position.X,
		player.Position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}
