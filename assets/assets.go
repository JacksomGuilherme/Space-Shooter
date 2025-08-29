package assets

import (
	"embed"
	"image"
	_ "image/png"
	"io/fs"
	"path"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed *
var assets embed.FS

// Space background
var BackgroundSprite = mustLoadImage("background.png")

// Player Sprites
var PlayerSpriteBlue = mustLoadImage("ships/ship_blue.png")
var PlayerSpriteRed = mustLoadImage("ships/ship_red.png")
var PlayerSpriteGreen = mustLoadImage("ships/ship_green.png")
var ShipShadowSprite = mustLoadImage("ships/ship_shadow.png")
var ShieldSprite = mustLoadImage("shield.png")

var PlatformSprite = mustLoadImage("platform.png")

// Entitys Sprites
var MeteorSprites = mustLoadMeteorsAssets("meteors/*.png")
var LaserSprite = mustLoadImage("laser.png")

// PowerUp Sprites
var HealingSprite = mustLoadImage("powerups/pill_green.png")
var BlueShieldSprite = mustLoadImage("powerups/blue_shield.png")
var BronzeStarSprite = mustLoadImage("powerups/star_bronze.png")
var SilverStarSprite = mustLoadImage("powerups/star_silver.png")
var GoldStarSprite = mustLoadImage("powerups/star_gold.png")

// Fonts
var ScoreFont = mustLoadFont("font.ttf")
var FontUi = mustLoadFont("fontui.ttf")

// Menu Sound Effects
var MenuSFX = mustLoadSFX("audio/SFX/beep.wav")
var MenuConfirmSFX = mustLoadSFX("audio/SFX/menu_confirm.wav")

// Player Sound Effects
var PlayerDeathSFX = mustLoadSFX("audio/SFX/player_death_whirl.wav")
var PlayerHitSFX = mustLoadSFX("audio/SFX/player_hit.wav")
var LaserSFX = mustLoadAllSFX("audio/SFX/laser/*.wav")
var ItemPickupSFX = mustLoadSFX("audio/SFX/item_pickup.wav")

type MeteorAsset struct {
	Image *ebiten.Image
	Color string
}

// Meteors Sound Effects
var MeteorsSFX = mustLoadAllSFX("audio/SFX/explosions/*.wav")

func mustLoadMeteorsAssets(pathPattern string) []*MeteorAsset {
	matches, err := fs.Glob(assets, pathPattern)
	if err != nil {
		panic(err)
	}
	meteorAssets := make([]*MeteorAsset, len(matches))
	for i, match := range matches {
		fileName := path.Base(match)
		color := "BROWN"
		if strings.Contains(fileName, "Grey") {
			color = "GREY"
		}
		meteorAssets[i] = &MeteorAsset{
			mustLoadImage(match),
			color,
		}
	}

	return meteorAssets
}

func mustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

// func mustLoadImages(path string) []*ebiten.Image {
// 	matches, err := fs.Glob(assets, path)
// 	if err != nil {
// 		panic(err)
// 	}

// 	images := make([]*ebiten.Image, len(matches))
// 	for i, match := range matches {
// 		images[i] = mustLoadImage(match)
// 	}

// 	return images
// }

func mustLoadFont(name string) font.Face {
	f, err := assets.ReadFile(name)
	if err != nil {
		panic(err)
	}

	tt, err := opentype.Parse(f)
	if err != nil {
		panic(err)
	}

	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    48,
		DPI:     72,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		panic(err)
	}

	return face
}

func GetFontFace(size int) font.Face {
	f, err := assets.ReadFile("fontui.ttf")
	if err != nil {
		panic(err)
	}

	tt, err := opentype.Parse(f)
	if err != nil {
		panic(err)
	}

	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(size),
		DPI:     72,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		panic(err)
	}

	return face
}

var audioContext = audio.NewContext(44100)

func mustLoadSFX(name string) []byte {
	data, err := assets.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return data
}

func mustLoadAllSFX(path string) [][]byte {
	matches, err := fs.Glob(assets, path)
	if err != nil {
		panic(err)
	}

	sounds := make([][]byte, len(matches))
	for i, match := range matches {
		data, err := assets.ReadFile(match)
		if err != nil {
			panic(err)
		}
		sounds[i] = data
	}
	return sounds
}

func PlaySFX(data []byte, volume float64) {
	p := audioContext.NewPlayerFromBytes(data)
	p.SetVolume(volume)
	p.Rewind()
	p.Play()
}
