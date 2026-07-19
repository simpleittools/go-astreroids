package assets

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image"
	_ "image/png"
	"io/fs"
)

//go:embed *
var assets embed.FS
var Explosion = createExplosion()
var ScoreFont = scoreFont("fonts/score.ttf")
var LevelFont = levelFont("fonts/score.ttf")
var TitleFont = titleFont("fonts/title.ttf")
var PlayerSprite = mustLoadImage("images/player.png")
var ShieldSprite = mustLoadImage("images/shield.png")
var MeteorSprites = mustLoadImages("images/meteors/*.png")
var MeteorSpritesSmall = mustLoadImages("images/meteors-small/*.png")
var AlienSprites = mustLoadImages("images/aliens/*.png")
var LaserSprite = mustLoadImage("images/laser.png")
var AlienLaserSprite = mustLoadImage("images/red-laser.png")
var ExhaustSprite = mustLoadImage("images/fire.png")
var ExplosionSprite = mustLoadImage("images/explosion.png")
var ExplosionSmallSprite = mustLoadImage("images/explosion-small.png")
var LaserOneSound = mustLoadOggVorbis("audio/fire.ogg")
var LaserTwoSound = mustLoadOggVorbis("audio/fire.ogg")
var LaserThreeSound = mustLoadOggVorbis("audio/fire.ogg")
var ThrustSound = mustLoadOggVorbis("audio/thrust.ogg")
var PlayerDiesSound = mustLoadOggVorbis("audio/player-dies.ogg")
var ExplosionSound = mustLoadOggVorbis("audio/explosion.ogg")
var AlienLaserSound = mustLoadOggVorbis("audio/alien-laser.ogg")
var ShieldSound = mustLoadOggVorbis("audio/shield.ogg")
var AlienSound = mustLoadOggVorbis("audio/alien-sound.ogg")
var BeatOneSound = mustLoadOggVorbis("audio/beat1.ogg")
var BeatTwoSound = mustLoadOggVorbis("audio/beat2.ogg")

var ShieldIndicator = mustLoadImage("images/shield-indicator.png")
var LifeIndicator = mustLoadImage("images/life-indicator.png")
var HyperspaceIndicator = mustLoadImage("images/hyperspace.png")

func createExplosion() []*ebiten.Image {
	var frames []*ebiten.Image
	for i := 0; i <= 11; i++ {
		frame := mustLoadImage(fmt.Sprintf("images/explosion/%d.png", i+1))
		frames = append(frames, frame)
	}
	return frames
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

func scoreFont(name string) font.Face {
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

func levelFont(name string) font.Face {
	f, err := assets.ReadFile(name)
	if err != nil {
		panic(err)
	}

	tt, err := opentype.Parse(f)
	if err != nil {
		panic(err)
	}

	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    12,
		DPI:     72,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		panic(err)
	}

	return face
}

func titleFont(name string) font.Face {
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
		Hinting: font.HintingFull,
	})
	if err != nil {
		panic(err)
	}

	return face
}

func mustLoadOggVorbis(name string) *vorbis.Stream {
	f, err := assets.ReadFile(name)
	if err != nil {
		panic(err)
	}
	d, err := vorbis.DecodeWithoutResampling(bytes.NewReader(f))
	if err != nil {
		panic(err)
	}
	return d
}
