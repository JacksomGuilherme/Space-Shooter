package game

import (
	"image/color"
	"math/rand"
	"space_shooter/assets"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// UpdateGameMode é responsável por atualizar a lógica do GameMode
func (game *Game) UpdateGameMode() {

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		game.viewport.moving = false
		game.NewPauseMode()
		game.mode = ModePause
	}

	game.Player.Update()

	newLasers := make([]*Laser, 0, len(game.Lasers))
	for _, laser := range game.Lasers {
		laser.Update()
		if laser.Position.Y+laser.Collider().Height > 0 {
			newLasers = append(newLasers, laser)
		}
	}
	game.Lasers = newLasers

	game.MeteorSpawnTimer.Update()
	if game.MeteorSpawnTimer.IsReady() {
		game.MeteorSpawnTimer.Reset()

		meteor := NewMeteor()

		game.Meteors = append(game.Meteors, meteor)
	}

	newMeteors := make([]*Meteor, 0, len(game.Meteors))
	for _, meteor := range game.Meteors {
		meteor.Update()
		if meteor.Position.Y < screenHeight {
			newMeteors = append(newMeteors, meteor)
		}
	}
	game.Meteors = newMeteors

	if game.Player.Ship.Health <= 0 {
		assets.PlaySound(game.Player.DeathSound, 1)
		game.mode = ModeGameOver
		game.gameOverCount = 30
		game.Reset()
	}

	for i, meteor := range game.Meteors {
		if meteor.Collider().Intersects(game.Player.Collider()) {
			if !meteor.Hit {
				meteor.Hit = true
				assets.PlaySound(game.Player.HitSound, 1)
				if !game.Player.Ship.ShieldActivated {
					game.Player.Ship.Health -= meteor.DamageByClass()
				}
				game.Meteors = append(game.Meteors[:i], game.Meteors[i+1:]...)
				break
			}
		}
	}

	game.PowerUpSpawnTimer.Update()
	for i := len(game.Meteors) - 1; i >= 0; i-- {
		meteor := game.Meteors[i]
		for j := len(game.Lasers) - 1; j >= 0; j-- {
			laser := game.Lasers[j]
			if (meteor.Position.Y + meteor.Collider().Height*0.9) > 0 {
				if meteor.Collider().Intersects(laser.Collider()) {
					assets.PlaySound(meteor.Sound, 1)
					if game.PowerUpSpawnTimer.IsReady() {
						game.PowerUpSpawnTimer.Reset()
						game.GeneratePowerUp(meteor)
					}
					game.Meteors = append(game.Meteors[:i], game.Meteors[i+1:]...)
					game.Lasers = append(game.Lasers[:j], game.Lasers[j+1:]...)
					game.Score++
					game.ScoreString = strconv.Itoa(game.Score)
					break
				}
			}
		}
	}

	newPowerUps := make([]*PowerUp, 0, len(game.PowerUps))
	for _, powerUp := range game.PowerUps {
		powerUp.Update()
		if powerUp.Position.Y < screenHeight {
			newPowerUps = append(newPowerUps, powerUp)
		}
	}
	game.PowerUps = newPowerUps

	for i, powerUp := range game.PowerUps {
		if powerUp.Collider().Intersects(game.Player.Collider()) {
			assets.PlaySound(powerUp.Sound, 1)
			powerUp.Action()
			game.PowerUps = append(game.PowerUps[:i], game.PowerUps[i+1:]...)
			break
		}
	}
}

// DrawGameMode é responsável por desenhar a tela de GameMode
func (game *Game) DrawGameMode(screen *ebiten.Image) {
	for _, laser := range game.Lasers {
		laser.Draw(screen)
	}

	game.Player.Draw(screen)

	for _, meteor := range game.Meteors {
		meteor.Draw(screen)
	}

	for _, powerUp := range game.PowerUps {
		powerUp.Draw(screen)
	}

	DrawHealthBar(screen, 20, screenHeight-40, 200, 20, game.Player.Ship.Health, 100)

	text.Draw(screen, "Points: "+game.ScoreString, ScoreFontFace, 20, 30, color.White)
}

// DrawHealthBar desenha a barra de vida na tela
func DrawHealthBar(screen *ebiten.Image, x, y, width, height float32, current, max int) {
	ratio := float32(current) / float32(max)
	filled := width * ratio

	barColor := color.RGBA{0, 200, 0, 255}
	shadowBarColor := color.RGBA{0, 180, 0, 200}
	if ratio < 0.5 {
		barColor = color.RGBA{200, 200, 0, 255}
		shadowBarColor = color.RGBA{180, 180, 0, 200}
	}
	if ratio < 0.25 {
		barColor = color.RGBA{200, 0, 0, 255}
		shadowBarColor = color.RGBA{180, 0, 0, 200}
	}

	text.Draw(screen, "HP", HealthFontFace, int(x), int(y-10), color.White)

	vector.DrawFilledRect(screen, x, y, width, height, color.RGBA{50, 50, 50, 255}, false)

	vector.DrawFilledRect(screen, x+2, y+2, width-4, height-4, color.RGBA{30, 30, 30, 255}, false)

	halfHeight := height / 2

	vector.DrawFilledRect(screen, x, y, filled, halfHeight, barColor, false)

	vector.DrawFilledRect(screen, x, y+halfHeight, filled, halfHeight, shadowBarColor, false)
}

func (game *Game) GeneratePowerUp(meteor *Meteor) {
	spawnPowerUp := rand.Intn(100) + 1
	if spawnPowerUp <= 20 && meteor.Color == "GREY" {
		meteorWidth := meteor.Collider().Width
		meteorHeight := meteor.Collider().Height

		position := Vector{
			meteor.Position.X + meteorWidth/2,
			meteor.Position.Y + meteorHeight/2,
		}

		powerUp := game.NewPowerUp(position)
		game.PowerUps = append(game.PowerUps, powerUp)
	}

}
