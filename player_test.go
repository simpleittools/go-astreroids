package main

import (
	"math"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	"simpleittools.com/go-asteroids/assets"
)

func TestNewPlayer(t *testing.T) {
	game := &Game{}

	player := NewPlayer(game)

	if player == nil {
		t.Fatal("expected a NewPlayer to return, but got nil")
	}

	if player.sprite == nil {
		t.Fatal("expected a NewPlayer to have a sprite, but got nil")
	}

	if player.sprite != assets.PlayerSprite {
		t.Fatal("expected a NewPlayer to have a sprite that is the same as the one in assets, but got a different one")
	}

	if player.rotation != 0 {
		t.Fatal("expected a NewPlayer to have a rotation of 0, but got a different one")
	}
}

func TestRotationPerSecond(t *testing.T) {
	if rotationPerSecond != math.Pi {
		t.Fatalf("expected rotationPerSecond to be Pi. Got %f", rotationPerSecond)
	}
}

func TestPlayerUpdateDoesNotRotateWhenNoKeysPressed(t *testing.T) {
	player := &Player{
		sprite: assets.PlayerSprite,
		rotation: 1.5,
	}

	player.Update()

	if player.rotation != 1.5 {
		t.Fatalf("expected player.rotation to be 1.5. Got %f", player.rotation)
	}
}

func TestPlayerDrawDoesNotPanic(t *testing.T) {
	player := &Player{
		sprite : assets.PlayerSprite,
		rotation: math.Pi / 2,
	}

	screen := ebiten.NewImage(320, 240)

	defer func() {
		if r := recover(); r!= nil {
			t.Fatalf("expected Draw to not panic. Got %v", r)
		}
	}()

	player.Draw(screen)
}