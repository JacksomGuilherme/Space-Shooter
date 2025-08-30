package assets

import (
	"embed"
	"image"
	_ "image/png"
	"io/fs"
	"path"
	"path/filepath"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed *
var assets embed.FS

// Map global para sons
var sounds = make(map[string][]byte)

// AudioContext
var audioContext = audio.NewContext(44100)

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

func init() {
	// Menu
	loadSound("menu_beep", "audio/SFX/beep.wav")
	loadSound("menu_confirm", "audio/SFX/menu_confirm.wav")

	// Player
	loadSound("player_death", "audio/SFX/player_death_whirl.wav")
	loadSound("player_hit", "audio/SFX/player_hit.wav")
	loadSound("item_pickup", "audio/SFX/item_pickup.wav")

	// Laser
	loadSound("laser", "audio/SFX/laser/laser_1.wav")

	// Explosões
	loadSound("explosion", "audio/SFX/explosions/explosion_1.wav")
}

type MeteorAsset struct {
	Image *ebiten.Image
	Color string
}

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

// ---------------- Funções de Som ----------------
func loadSound(name, path string) {
	data, err := assets.ReadFile(path)
	if err != nil {
		panic(err)
	}
	sounds[name] = data
}

func loadAllSounds(prefix, pathPattern string) {
	matches, err := fs.Glob(assets, pathPattern)
	if err != nil {
		panic(err)
	}

	for _, match := range matches {
		data, err := assets.ReadFile(match)
		if err != nil {
			panic(err)
		}
		key := prefix + strings.TrimSuffix(filepath.Base(match), filepath.Ext(match))
		sounds[key] = data
	}
}

// Função pública para tocar som pelo nome
func PlaySound(name string, volume float64) {
	data, ok := sounds[name]
	if !ok {
		return
	}
	p := audioContext.NewPlayerFromBytes(data)
	p.SetVolume(volume)
	p.Rewind()
	p.Play()
}
