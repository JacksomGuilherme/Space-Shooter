package game

import (
	"image/color"
	"space_shooter/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type Menu struct {
	Items             []MenuItem
	Images            []MenuItemImage
	SelectedItemIndex int
}

type MenuItem struct {
	Label  string
	Action func()
}

type MenuItemImage struct {
	Image  *ebiten.Image
	Action func()
}

func NewMenu(items []MenuItem) *Menu {
	return &Menu{
		Items: items,
	}
}

func NewMenuImages(images []MenuItemImage) *Menu {
	return &Menu{
		Images:            images,
		SelectedItemIndex: 0,
	}
}

func (game *Game) UpdateMenuText() {
	menuSound := assets.MenuSFX
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		selectedIndex := game.Menu.SelectedItemIndex
		assets.PlaySFX(assets.MenuConfirmSFX, 1)
		game.Menu.Items[selectedIndex].Action()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		assets.PlaySFX(menuSound, 1)

		if game.Menu.SelectedItemIndex == len(game.Menu.Items)-1 {
			return
		}

		game.Menu.SelectedItemIndex += 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		assets.PlaySFX(menuSound, 1)

		if (game.Menu.SelectedItemIndex - 1) < 0 {
			return
		}

		game.Menu.SelectedItemIndex -= 1
	}
}

func (game *Game) UpdateMenuImage() {
	menuSound := assets.MenuSFX
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		selectedIndex := game.Menu.SelectedItemIndex
		assets.PlaySFX(assets.MenuConfirmSFX, 1)
		game.Menu.Images[selectedIndex].Action()
		game.mode = ModeTitle
		game.NewTitleMode()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		assets.PlaySFX(menuSound, 1)

		if (game.Menu.SelectedItemIndex - 1) < 0 {
			return
		}

		game.Menu.SelectedItemIndex -= 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		assets.PlaySFX(menuSound, 1)

		if game.Menu.SelectedItemIndex == len(game.Menu.Images)-1 {
			return
		}

		game.Menu.SelectedItemIndex += 1
	}
}

func (game *Game) DrawMenuText(screen *ebiten.Image, initalPos int) {
	verticalMenuItemPos := initalPos

	for i, menu := range game.Menu.Items {
		colorMenuItem := color.RGBA{255, 255, 255, 255}

		if game.Menu.SelectedItemIndex == i {
			colorMenuItem = color.RGBA{102, 255, 255, 255}
		}

		fontFace := assets.GetFontFace(32)
		texts := menu.Label
		bounds, _ := font.BoundString(fontFace, texts)
		textHalfWidth := bounds.Max.X / 2
		text.Draw(screen, texts, fontFace, (screenWidth/2 - textHalfWidth.Round()), verticalMenuItemPos, colorMenuItem)
		verticalMenuItemPos += 40
	}
	game.Player.Draw(screen)
}

func (game *Game) DrawMenuImage(screen *ebiten.Image) {
	for i, menu := range game.Menu.Images {
		if i == game.Menu.SelectedItemIndex {
			bounds := menu.Image.Bounds()
			halfW := (float64(bounds.Dx()) * 2.1) / 2
			halfY := (float64(bounds.Dy()) * 2.1) / 2

			shipShadow := assets.ShipShadowSprite

			options := &ebiten.DrawImageOptions{}
			options.GeoM.Scale(2.1, 2.1)
			options.GeoM.Translate(screenWidth/2-halfW, screenHeight/2-halfY)

			screen.DrawImage(shipShadow, options)

			halfW = (float64(bounds.Dx()) * 2) / 2
			halfY = (float64(bounds.Dy()) * 2) / 2

			options = &ebiten.DrawImageOptions{}
			options.GeoM.Scale(2, 2)
			options.GeoM.Translate(screenWidth/2-halfW, screenHeight/2-halfY)

			screen.DrawImage(menu.Image, options)
		}
	}
}
