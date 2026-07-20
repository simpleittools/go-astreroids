package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"simpleittools.com/go-asteroids/assets"
)

const (
	rotationPerSecond = math.Pi
	maxAcceleration   = 8.0
	// ScreenWidth sets the screen width
	ScreenWidth = 1280
	// ScreenHeight sets the screen height. This keeps a 16:9 aspect ratio
	ScreenHeight = 720
)

// curAcceleration is how fast the player is going, and it will increase or decrease over time.
var curAcceleration float64

var (
	isKeyPressed = ebiten.IsKeyPressed
	tps          = ebiten.TPS
)

type Player struct {
	sprite *ebiten.Image
	// what is the rotation of the player sprite
	rotation       float64
	game           *Game
	position       Vector
	playerVelocity float64
}

func NewPlayer(game *Game) *Player {
	sprite := assets.PlayerSprite

	// center player on the screen
	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	pos := Vector{
		X: float64(ScreenWidth) / 2 - halfW,
		Y: float64(ScreenHeight) / 2 - halfH}

	p := &Player{
		sprite: sprite,
		// by adding game, it gives the player access to items outside itself and part of the overall game
		game: game,
		position: pos,
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

	// move the player to the position if necessary
	op.GeoM.Translate(p.position.X, p.position.Y)

	screen.DrawImage(p.sprite, op)

}

// Update updates the player and must be called Update.
func (p *Player) Update() {
	// We want to rotate the image on the screen. ebiten.TPS is ticks per second. This determines how fast the player rotates.s
	speed := rotationPerSecond / float64(tps())

	// rotate the player sprite when the player presses the left arrow key
	if isKeyPressed(ebiten.KeyA) {
		p.rotation -= speed
	}

	if isKeyPressed(ebiten.KeyD) {
		p.rotation += speed
	}

	p.accelerate()
}

func (p *Player) accelerate() {
	if isKeyPressed(ebiten.KeyW) {
		p.KeepOnScreen()


		// perform a gradual increase in acceleration
		if curAcceleration < maxAcceleration {
			curAcceleration = p.playerVelocity + 4
		}

		if curAcceleration >= 8 {
			curAcceleration = 8
		}

		p.playerVelocity = curAcceleration

		// Move in the direction we are pointing
		dx := math.Sin(p.rotation) * curAcceleration
		dy := math.Cos(p.rotation) * -curAcceleration

		// Move the player on the screen
		p.position.X += dx
		p.position.Y += dy
	}
}

// KeepOnScreen keeps the player on the screen.
func (p *Player) KeepOnScreen() {
	if p.position.X >= float64(ScreenWidth) {
		p.position.X = 0
	}

	if p.position.X < 0 {
		p.position.X = float64(ScreenWidth)
	}

	if p.position.Y >= float64(ScreenHeight) {
		p.position.Y = 0
	}

	if p.position.Y < 0 {
		p.position.Y = float64(ScreenHeight)
	}
}
