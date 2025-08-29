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
	SelectedItemIndex int
}

type MenuItem struct {
	Label  string
	Action func()
}

func NewMenu() *Menu {
	return &Menu{
		Items: nil,
	}
}

func (game *Game) DrawMenu(screen *ebiten.Image) {
	verticalMenuItemPos := int(screenHeight - titleFontSize*8)

	for i, menu := range game.Menu.Items {
		colorMenuItem := color.RGBA{255, 255, 255, 255}

		if game.Menu.SelectedItemIndex == i {
			colorMenuItem = color.RGBA{0, 200, 0, 255}
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

func (game *Game) UpdateMenu() {
	menuSound := assets.MenuSFX
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		selectedIndex := game.Menu.SelectedItemIndex
		game.Menu.Items[selectedIndex].Action()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		assets.PlaySFX(menuSound, 1)

		if game.Menu.SelectedItemIndex == len(game.Menu.Items)-1 {
			return
		}

		game.Menu.SelectedItemIndex += 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		assets.PlaySFX(menuSound, 1)

		if (game.Menu.SelectedItemIndex - 1) < 0 {
			return
		}

		game.Menu.SelectedItemIndex -= 1
	}
}
