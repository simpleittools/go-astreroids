package main

import (
	"math"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	"simpleittools.com/go-asteroids/assets"
)

func withPlayerInputStub(t *testing.T, pressed map[ebiten.Key]bool, ticksPerSecond int) {
	t.Helper()

	originalIsKeyPressed := isKeyPressed
	originalTPS := tps
	originalAcceleration := curAcceleration

	isKeyPressed = func(key ebiten.Key) bool {
		return pressed[key]
	}
	tps = func() int {
		return ticksPerSecond
	}
	curAcceleration = 0

	t.Cleanup(func() {
		isKeyPressed = originalIsKeyPressed
		tps = originalTPS
		curAcceleration = originalAcceleration
	})
}

func TestNewPlayer(t *testing.T) {
	game := &Game{}

	player := NewPlayer(game)

	if player == nil {
		t.Fatal("expected NewPlayer to return a player, got nil")
	}

	if player.sprite == nil {
		t.Fatal("expected player sprite to be initialized")
	}

	if player.sprite != assets.PlayerSprite {
		t.Fatal("expected player sprite to match assets.PlayerSprite")
	}

	if player.rotation != 0 {
		t.Fatalf("expected rotation to start at 0, got %f", player.rotation)
	}
}

func TestRotationPerSecond(t *testing.T) {
	if rotationPerSecond != math.Pi {
		t.Fatalf("expected rotationPerSecond to be Pi, got %f", rotationPerSecond)
	}
}

func TestPlayerUpdateDoesNotRotateWhenNoKeysPressed(t *testing.T) {
	withPlayerInputStub(t, map[ebiten.Key]bool{}, 60)

	player := &Player{
		sprite:   assets.PlayerSprite,
		rotation: 1.5,
	}

	player.Update()

	if player.rotation != 1.5 {
		t.Fatalf("expected rotation to remain 1.5, got %f", player.rotation)
	}
}

func TestPlayerUpdateRotatesLeftAndRight(t *testing.T) {
	tests := []struct {
		name     string
		pressed  map[ebiten.Key]bool
		expected float64
	}{
		{
			name:     "left",
			pressed:  map[ebiten.Key]bool{ebiten.KeyLeft: true},
			expected: -math.Pi / 60,
		},
		{
			name:     "right",
			pressed:  map[ebiten.Key]bool{ebiten.KeyRight: true},
			expected: math.Pi / 60,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			withPlayerInputStub(t, tc.pressed, 60)

			player := &Player{sprite: assets.PlayerSprite}
			player.Update()

			if math.Abs(player.rotation-tc.expected) > 1e-9 {
				t.Fatalf("expected rotation %f, got %f", tc.expected, player.rotation)
			}
		})
	}
}

func TestPlayerUpdateAcceleratesAndMovesForward(t *testing.T) {
	withPlayerInputStub(t, map[ebiten.Key]bool{ebiten.KeyUp: true}, 60)

	player := &Player{
		sprite:   assets.PlayerSprite,
		rotation: 0,
	}

	player.Update()

	if player.playerVelocity != 4 {
		t.Fatalf("expected playerVelocity to be 4, got %f", player.playerVelocity)
	}

	if curAcceleration != 4 {
		t.Fatalf("expected curAcceleration to be 4, got %f", curAcceleration)
	}

	if math.Abs(player.position.X) > 1e-9 {
		t.Fatalf("expected X position to remain 0, got %f", player.position.X)
	}

	if math.Abs(player.position.Y+4) > 1e-9 {
		t.Fatalf("expected Y position to move by -4, got %f", player.position.Y)
	}
}

func TestPlayerUpdateCapsAccelerationAtMax(t *testing.T) {
	withPlayerInputStub(t, map[ebiten.Key]bool{ebiten.KeyUp: true}, 60)

	player := &Player{
		sprite:   assets.PlayerSprite,
		rotation: 0,
	}

	player.Update()
	player.Update()
	player.Update()

	if player.playerVelocity != maxAcceleration {
		t.Fatalf("expected playerVelocity to cap at %f, got %f", maxAcceleration, player.playerVelocity)
	}

	if curAcceleration != maxAcceleration {
		t.Fatalf("expected curAcceleration to cap at %f, got %f", maxAcceleration, curAcceleration)
	}

	if math.Abs(player.position.Y+16) > 1e-9 {
		t.Fatalf("expected cumulative Y movement to be -16, got %f", player.position.Y)
	}
}

func TestPlayerDrawDoesNotPanic(t *testing.T) {
	player := &Player{
		sprite:   assets.PlayerSprite,
		rotation: math.Pi / 2,
	}

	screen := ebiten.NewImage(320, 240)

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("expected Draw to not panic, got %v", r)
		}
	}()

	player.Draw(screen)
}
