package game

import (
	"image/color"
	"space_shooter/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

// UpdateGameOverMode é responsável por atualizar a lógica da tela de Game Over
func (game *Game) UpdateGameOverMode() {
	if game.gameOverCount > 0 {
		game.gameOverCount--
	}
	if game.gameOverCount == 0 && inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		menuSound := assets.MenuSFX
		game.Reset()
		assets.PlaySFX(menuSound, 1)
		game.mode = ModeTitle
	}
}

// DrawGameOverMode é responsável desenhar a tela de Game Over
func (game *Game) DrawGameOverMode(screen *ebiten.Image) {
	titleText := "GAME OVER!"
	bounds, _ := font.BoundString(assets.FontUi, titleText)
	textHalfWidth := bounds.Max.X / 2
	textHalfHeight := bounds.Max.Y / 2
	text.Draw(
		screen,
		titleText,
		assets.FontUi,
		(screenWidth/2 - textHalfWidth.Round()),
		(screenHeight/2 - textHalfHeight.Round()),
		color.White)
}
