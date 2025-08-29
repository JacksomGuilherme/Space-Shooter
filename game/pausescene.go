package game

import (
	"image/color"
	"space_shooter/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

// UpdatePauseMode é responsável por atualizar a lógca da tela de pause
func (game *Game) UpdatePauseMode() {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		game.viewport.moving = true
		game.mode = ModeGame
	}
}

// DrawPauseMode é responsável por desenhar a tela de pausa
func (game *Game) DrawPauseMode(screen *ebiten.Image) {
	titleText := "PAUSE"
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

	fontFace := assets.GetFontFace(32)
	texts := "PRESS SPACE TO CONTINUE"
	bounds, _ = font.BoundString(fontFace, texts)
	textHalfWidth = bounds.Max.X / 2
	text.Draw(screen, texts, fontFace, (screenWidth/2 - textHalfWidth.Round()), screenHeight-titleFontSize*5, color.White)
}
