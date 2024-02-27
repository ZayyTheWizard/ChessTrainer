package gamemain

import (
	"log"

	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	board BoardDimension
}

var (
	board = BoardDimension{
		ROW:      8,
		COL:      8,
		rowColor: color.RGBA{R: 101, G: 67, B: 33, A: 255},
		colColor: color.RGBA{R: 255, G: 239, B: 215, A: 255},
		consoleRepresentation: [64]PieceDescription{
			{PieceName: Rook, color: "black"}, {PieceName: Knight, color: "black"}, {PieceName: Bishop, color: "black"}, {PieceName: Queen, color: "black"}, {PieceName: King, color: "black"}, {PieceName: Bishop, color: "black"}, {PieceName: Knight, color: "black"}, {PieceName: Rook, color: "black"},
			{PieceName: Pawn, color: "black"}, {PieceName: Pawn, color: "black"}, {PieceName: Pawn, color: "black"}, {PieceName: Pawn, color: "black"}, {PieceName: Pawn, color: "black"}, {PieceName: Pawn, color: "black"}, {PieceName: Pawn, color: "black"}, {PieceName: Pawn, color: "black"},
			{}, {}, {}, {}, {}, {}, {}, {},
			{}, {}, {}, {}, {}, {}, {}, {},
			{}, {}, {}, {}, {}, {}, {}, {},
			{}, {}, {}, {}, {}, {}, {}, {},
			{PieceName: Pawn, color: "white"}, {PieceName: Pawn, color: "white"}, {PieceName: Pawn, color: "white"}, {PieceName: Pawn, color: "white"}, {PieceName: Pawn, color: "white"}, {PieceName: Pawn, color: "white"}, {PieceName: Pawn, color: "white"}, {PieceName: Pawn, color: "white"},
			{PieceName: Rook, color: "white"}, {PieceName: Knight, color: "white"}, {PieceName: Bishop, color: "white"}, {PieceName: Queen, color: "white"}, {PieceName: King, color: "white"}, {PieceName: Bishop, color: "white"}, {PieceName: Knight, color: "white"}, {PieceName: Rook, color: "white"},
		},
	}
)

func (g *Game) Update() error {

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	board.drawBoard(screen)
}

func (g *Game) Layout(OutsideWidth, OutsideHeight int) (screenWidth, screenHeight int) {
	return 1000, 1000
}

func Start() {
	ebiten.SetWindowSize(1280, 1280)
	ebiten.SetWindowTitle("Chess")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
