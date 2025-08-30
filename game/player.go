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
	DeathSound        string
	HitSound          string
	LaserLoadingTimer *Timer
}

type Ship struct {
	Image           *ebiten.Image
	ShieldImage     *ebiten.Image
	Health          int
	Laser           *ebiten.Image
	ShieldActivated bool
	ShieldTimer     *Timer
}

var (
	speed        = 6.0
	playerBounds = assets.PlayerSpriteBlue.Bounds()
	playerHalfW  = float64(playerBounds.Dx()) / 2
	playerHalfH  = float64(playerBounds.Dy()) / 2

	shipBounds = assets.ShieldSprite.Bounds()
	shipHalfW  = float64(shipBounds.Dx()) / 2
	shipHalfH  = float64(shipBounds.Dy()) / 2
)

// NewPlayer é responsável por criar uma instância de Player
func NewPlayer(game *Game) *Player {
	ship := Ship{
		Image:       assets.PlayerSpriteBlue,
		ShieldImage: assets.ShieldSprite,
		Health:      100,
		ShieldTimer: NewTimer(600),
	}

	if game.Player != nil && game.Player.Ship.Image != nil {
		ship.Image = game.Player.Ship.Image
	}

	position := Vector{
		X: screenWidth/2 - playerHalfW,
		Y: 500,
	}

	deathSound := "player_death"
	hitSound := "player_hit"

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
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if (player.Position.X - speed) <= 0-playerHalfW {
			return
		}

		player.Position.X -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if (player.Position.X + speed) >= screenWidth-playerHalfW {
			return
		}
		player.Position.X += speed
	}

	player.LaserLoadingTimer.Update()
	if ebiten.IsKeyPressed(ebiten.KeySpace) && player.LaserLoadingTimer.IsReady() {
		player.LaserLoadingTimer.Reset()

		spawnPosition := Vector{
			player.Position.X + playerHalfW,
			player.Position.Y - playerHalfH/2,
		}
		laser := NewLaser(spawnPosition)
		assets.PlaySound(laser.Sound, 1)
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
		options = &ebiten.DrawImageOptions{}

		playerCenterX := player.Position.X + playerHalfW
		playerCenterY := player.Position.Y + playerHalfH

		shieldX := playerCenterX - shipHalfW
		shieldY := playerCenterY - shipHalfH

		options.GeoM.Translate(shieldX, shieldY)
		screen.DrawImage(player.Ship.ShieldImage, options)
	}

}

// Collider determina as dimensões do retângulo de hitbox do player
func (player *Player) Collider() Rect {
	return NewRect(
		player.Position.X,
		player.Position.Y,
		float64(playerBounds.Dx()),
		float64(playerBounds.Dy()),
	)
}
