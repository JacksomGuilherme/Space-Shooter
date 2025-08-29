package game

import (
	"image/color"
	"space_shooter/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

func (game *Game) NewShipSelectionMode() {
	game.Menu = NewMenuImages([]MenuItemImage{
		{assets.PlayerSpriteBlue, func() { game.Player.Ship.Image = assets.PlayerSpriteBlue }},
		{assets.PlayerSpriteRed, func() { game.Player.Ship.Image = assets.PlayerSpriteRed }},
		{assets.PlayerSpriteGreen, func() { game.Player.Ship.Image = assets.PlayerSpriteGreen }},
	})
}

// UpdateShipSelectionMode é responsável por atualizar a lógca da tela de seleção de nave
func (game *Game) UpdateShipSelectionMode() {
	game.UpdateMenuImage()
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		game.NewTitleMode()
		game.mode = ModeTitle
	}
}

// DrawShipSelectionMode é responsável por desenhar a tela de seleção de nave
func (game *Game) DrawShipSelectionMode(screen *ebiten.Image) {
	platform := assets.PlatformSprit
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(platform, op)

	titleText := "Ship Selection"
	bounds, _ := font.BoundString(assets.FontUi, titleText)
	textHalfWidth := bounds.Max.X / 2
	text.Draw(screen, titleText, assets.FontUi, (screenWidth/2 - textHalfWidth.Round()), 3*titleFontSize, color.White)

	game.DrawMenuImage(screen)

	textInitialVerticalPos := screenHeight * 0.85

	fontFace := assets.GetFontFace(32)
	texts := "PRESS ENTER TO SELECT"
	bounds, _ = font.BoundString(fontFace, texts)
	textHalfWidth = bounds.Max.X / 2
	text.Draw(screen, texts, fontFace, (screenWidth/2 - textHalfWidth.Round()), int(textInitialVerticalPos), color.White)

	texts = "PRESS ESC TO GO BACK"
	bounds, _ = font.BoundString(fontFace, texts)
	textHalfWidth = bounds.Max.X / 2
	text.Draw(screen, texts, fontFace, (screenWidth/2 - textHalfWidth.Round()), int(textInitialVerticalPos+40), color.White)
}
