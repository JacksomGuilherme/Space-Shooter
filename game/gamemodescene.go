package game

import (
	"fmt"
	"image/color"
	"space_shooter/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// UpdateGameMode é responsável por atualizar a lógica do GameMode
func (game *Game) UpdateGameMode() {

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		game.viewport.moving = false
		game.mode = ModePause
	}

	game.Player.Update()

	for _, laser := range game.Lasers {
		laser.Update()
	}

	game.MeteorSpawnTimer.Update()
	if game.MeteorSpawnTimer.IsReady() {
		game.MeteorSpawnTimer.Reset()

		meteor := NewMeteor()

		game.Meteors = append(game.Meteors, meteor)
	}

	for _, meteor := range game.Meteors {
		meteor.Update()
	}

	if game.Player.Health <= 0 {
		assets.PlaySFX(game.Player.DeathSound, 1)
		game.mode = ModeGameOver
		game.gameOverCount = 30
		game.Reset()
	}

	for i, meteor := range game.Meteors {
		if meteor.Collider().Intersects(game.Player.Collider()) {
			if !meteor.Hit {
				assets.PlaySFX(game.Player.HitSound, 1)
				game.Player.Health -= meteor.DamageByClass()
				meteor.Hit = true
				game.Meteors = append(game.Meteors[:i], game.Meteors[i+1:]...)
				break
			}
		}
	}

	for i, meteor := range game.Meteors {
		for j, laser := range game.Lasers {
			if meteor.Collider().Intersects(laser.Collider()) {
				assets.PlaySFX(meteor.Sound, 1)
				game.Meteors = append(game.Meteors[:i], game.Meteors[i+1:]...)
				game.Lasers = append(game.Lasers[:j], game.Lasers[j+1:]...)
				game.Score++
				break
			}
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

	DrawHealthBar(screen, 20, screenHeight-40, 200, 20, game.Player.Health, 100)

	text.Draw(screen, fmt.Sprintf("Points: %d", game.Score), assets.GetFontFace(24), 20, 30, color.White)
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

	text.Draw(screen, "HP", assets.GetFontFace(24), int(x), int(y-10), color.White)

	vector.DrawFilledRect(screen, x, y, width, height, color.RGBA{50, 50, 50, 255}, false)

	vector.DrawFilledRect(screen, x+2, y+2, width-4, height-4, color.RGBA{30, 30, 30, 255}, false)

	halfHeight := height / 2

	vector.DrawFilledRect(screen, x, y, filled, halfHeight, barColor, false)

	vector.DrawFilledRect(screen, x, y+halfHeight, filled, halfHeight, shadowBarColor, false)
}
