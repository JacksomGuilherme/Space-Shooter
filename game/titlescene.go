package game

import (
	"image/color"
	"space_shooter/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

// UpdateTitleMode é responsável por atualizar a lógica de Title
func (game *Game) UpdateTitleMode() {
	game.UpdateMenu()
}

// DrawTitleMode é responsável por desenhar a tela de Title
func (game *Game) DrawTitleMode(screen *ebiten.Image) {
	game.Player.Draw(screen)

	titleText := "Space Shooter"
	bounds, _ := font.BoundString(assets.FontUi, titleText)
	textHalfWidth := bounds.Max.X / 2
	text.Draw(screen, titleText, assets.FontUi, (screenWidth/2 - textHalfWidth.Round()), 3*titleFontSize, color.White)

	game.DrawMenu(screen)
}
