package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"simpleittools.com/go-asteroids/assets"
)

type Player struct {
	sprite *ebiten.Image
	// what is the rotation of the player sprite
	rotation float64
}

const rotationPerSecond = math.Pi

func NewPlayer(game *Game) *Player {
	sprite := assets.PlayerSprite

	p := &Player{
		sprite: sprite,
	}

	return p
}

// Draw draws the player and must be called Draw. It must take the argument of screen from *ebiten.Image
func (p *Player) Draw(screen *ebiten.Image) {
	// Bounds gives us the rectangle that is drawn over the image. Basically, what are the boundaries of the player image.
	bounds := p.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	// options
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(p.rotation)
	op.GeoM.Translate(halfW, halfH)

	screen.DrawImage(p.sprite, op)

}

// Update updates the player and must be called Update.
func (p *Player) Update() {
	// We want to rotate the image on the screen. ebiten.TPS is ticks per second. This determines how fast the player rotates.s
	speed := rotationPerSecond / float64(ebiten.TPS())

	// rotate the player sprite when the player presses the left arrow key
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.rotation -= speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.rotation += speed
	}
}