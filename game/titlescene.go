package game

import (
	"fmt"
	"image/color"
	"os"
	"space_shooter/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

func (game *Game) NewTitleMode() {
	game.Menu = NewMenu([]MenuItem{
		{"NEW GAME", func() { game.mode = ModeGame }},
		{"CHOOSE SHIP", func() {
			game.NewShipSelectionMode()
			game.mode = ModeShipSelection
		}},
		{"EXIT GAME", func() { os.Exit(0) }},
	})
}

// UpdateTitleMode é responsável por atualizar a lógica de Title
func (game *Game) UpdateTitleMode() {
	game.UpdateMenuText()
}

// DrawTitleMode é responsável por desenhar a tela de Title
func (game *Game) DrawTitleMode(screen *ebiten.Image) {
	game.Player.Draw(screen)

	titleText := "Space Shooter"
	bounds, _ := font.BoundString(assets.FontUi, titleText)
	textHalfWidth := bounds.Max.X / 2
	text.Draw(screen, titleText, assets.FontUi, (screenWidth/2 - textHalfWidth.Round()), 3*titleFontSize, color.White)

	fontFace := assets.GetFontFace(32)
	texts := fmt.Sprintf("Max Score: %d", game.MaxScore)
	bounds, _ = font.BoundString(fontFace, texts)
	textHalfWidth = bounds.Max.X / 2
	text.Draw(screen, texts, fontFace, (screenWidth/2 - textHalfWidth.Round()), 3*titleFontSize+40, color.White)

	game.DrawMenuText(screen, int(screenHeight-titleFontSize*8))
}
