package assets

import (
	"embed"
	"image"
	_ "image/png"
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed *
var assets embed.FS

var PlayerSprite = mustLoadImage("player.png")
var BackgroundSprite = mustLoadImage("background.png")

var MeteorSprites = mustLoadImages("meteors/*.png")
var LaserSprite = mustLoadImage("laser.png")
var GopherPlayer = mustLoadImage("go_player.png")
var StarsSprites = mustLoadImages("stars/*.png")
var PlanetsSprites = mustLoadImages("planets/*.png")

var ScoreFont = mustLoadFont("font.ttf")
var FontUi = mustLoadFont("fontui.ttf")

var MenuSFX = mustLoadSFX("audio/SFX/beep.wav")

var PlayerDeathSFX = mustLoadSFX("audio/SFX/player_death_whirl.wav")
var PlayerHitSFX = mustLoadSFX("audio/SFX/player_hit.wav")
var LaserSFX = mustLoadAllSFX("audio/SFX/laser/*.wav")

var MeteorsSFX = mustLoadAllSFX("audio/SFX/explosions/*.wav")

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

func mustLoadImages(path string) []*ebiten.Image {
	matches, err := fs.Glob(assets, path)
	if err != nil {
		panic(err)
	}

	images := make([]*ebiten.Image, len(matches))
	for i, match := range matches {
		images[i] = mustLoadImage(match)
	}

	return images
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
