package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"simpleittools.com/go-asteroids/assets"
)

type Player struct {
	sprite *ebiten.Image
}

func NewPlayer(game *Game) *Player {
	sprite := assets.PlayerSprite

	p := &Player{
		sprite: sprite,
	}

	return p
}

// Draw draws the player and must be called Draw. It must take the argument of screen from *ebiten.Image
func (p *Player) Draw(screen *ebiten.Image) {
	// options
	op := &ebiten.DrawImageOptions{
		// TODO: populate
	}

	screen.DrawImage(p.sprite, op)

}

// Update updates the player and must be called Update.
func (p *Player) Update() {

}