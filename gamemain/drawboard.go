package gamemain

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type BoardDimension struct {
	ROW                   int
	COL                   int
	rowColor              color.Color
	colColor              color.Color
	consoleRepresentation [64]PieceDescription
}

func (inf *BoardDimension) updateBoard() error {

	return nil
}

func (inf BoardDimension) drawBoard(screen *ebiten.Image) {
	screenWidth, screenHeight := screen.Bounds().Dx(), screen.Bounds().Dy()
	squareSize := int(math.Min(float64(screenWidth)/float64(inf.COL), float64(screenHeight)/float64(inf.ROW)))

	rowSq := ebiten.NewImage(squareSize, squareSize)
	colSq := ebiten.NewImage(squareSize, squareSize)

	rowSq.Fill(inf.rowColor)
	colSq.Fill(inf.colColor)

	opts := &ebiten.DrawImageOptions{}

	pieceImageMap := initImages()
	for i := 0; i < inf.ROW; i++ {
		for j := 0; j < inf.COL; j++ {
			opts.GeoM.Reset()
			opts.GeoM.Translate(float64(j*squareSize), float64(i*squareSize))

			if (i+j)%2 == 0 {
				screen.DrawImage(colSq, opts)
			} else {
				screen.DrawImage(rowSq, opts)
			}

			// Draw the pieces
			pieceDis := inf.consoleRepresentation[i*inf.COL+j]
			if pieceDis.PieceName != 0 {
				opts := &ebiten.DrawImageOptions{}
				opts.GeoM.Scale(2, 2)
				opts.GeoM.Translate(float64(j*squareSize), float64(i*squareSize))
				screen.DrawImage(pieceImageMap[pieceDis], opts)
			}

		}
	}

	row, col := inf.mouseLoc(screen)
	inf.highlightPiece(screen, row, col)
}
