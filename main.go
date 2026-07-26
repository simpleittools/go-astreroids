package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"simpleittools.com/go-asteroids/goasteroids"
)

type Game struct {
	player *goasteroids.Player

}

func (g *Game) Update() error {
	g.player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (ScreenWidth, ScreenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {

	g := &Game{

	}
	ebiten.SetWindowTitle("Go Asteroids")
	ebiten.SetWindowSize(goasteroids.ScreenWidth, goasteroids.ScreenHeight)

	g.player = goasteroids.NewPlayer(g)

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}