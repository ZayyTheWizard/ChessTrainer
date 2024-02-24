package gamemain

import (
	"log"

	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	board BoardDimension
}

var (
	board = BoardDimension{
		ROW: 8,
		COL: 8,
	}
)

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello World!")
	board.drawSquars(screen, color.RGBA{R: 255, G: 0, B: 0, A: 255}, color.RGBA{R: 0, G: 255, B: 0, A: 255})
}

func (g *Game) Layout(OutsideWidth, OutsideHeight int) (screenWidth, screenHeight int) {
	return 1000, 1000
}

func Start() {
	ebiten.SetWindowSize(1280, 1280)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
