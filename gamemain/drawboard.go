package gamemain

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type BoardDimension struct {
	ROW int
	COL int
}

func (inf BoardDimension) drawSquars(screen *ebiten.Image, rowColor, colColor color.Color) {
	screenWidth, screenHeight := screen.Size()
	squareSize := int(math.Min(float64(screenWidth)/float64(inf.COL), float64(screenHeight)/float64(inf.ROW)))

	rowSq := ebiten.NewImage(squareSize, squareSize)
	colSq := ebiten.NewImage(squareSize, squareSize)

	rowSq.Fill(rowColor)
	colSq.Fill(colColor)

	opts := &ebiten.DrawImageOptions{}

	for i := 0; i < inf.ROW; i++ {
		for j := 0; j < inf.COL; j++ {
			opts.GeoM.Reset()
			opts.GeoM.Translate(float64(j*squareSize), float64(i*squareSize))

			if (i+j)%2 == 0 {
				screen.DrawImage(colSq, opts)
			} else {
				screen.DrawImage(rowSq, opts)
			}
		}
	}
}
