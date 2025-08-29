package game

import (
	"image/color"
	"space_shooter/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

func (game *Game) NewPauseMode() {
	game.Menu = NewMenu([]MenuItem{
		{"RESUME", func() { game.mode = ModeGame }},
		{"MAIN MENU", func() {
			game.Reset()
			game.NewTitleMode()
			game.mode = ModeTitle
		}},
	})
}

// UpdatePauseMode é responsável por atualizar a lógca da tela de pause
func (game *Game) UpdatePauseMode() {
	game.UpdateMenuText()
}

// DrawPauseMode é responsável por desenhar a tela de pausa
func (game *Game) DrawPauseMode(screen *ebiten.Image) {
	titleText := "PAUSE"
	bounds, _ := font.BoundString(assets.FontUi, titleText)
	textHalfWidth := bounds.Max.X / 2
	text.Draw(
		screen,
		titleText,
		assets.FontUi,
		(screenWidth/2 - textHalfWidth.Round()),
		3*titleFontSize,
		color.White)

	game.DrawMenuText(screen, int(screenHeight-titleFontSize*8))
}
