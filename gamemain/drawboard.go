package gamemain

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type BoardDimension struct {
	ROW                   int
	COL                   int
	rowColor              color.Color
	colColor              color.Color
	consoleRepresentation [64]PieceDescription
}

func (inf BoardDimension) drawSquares(screen *ebiten.Image) {
	screenWidth, screenHeight := screen.Size()
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
}

// load piece images here
func initImages() map[PieceDescription]*ebiten.Image {
	PawnW, _, err := ebitenutil.NewImageFromFile("./ChessImages/pawn_white.png")
	ThrowEror(err)
	PawnB, _, _ := ebitenutil.NewImageFromFile("./ChessImages/pawn_black.png")
	ThrowEror(err)
	BishopW, _, _ := ebitenutil.NewImageFromFile("./ChessImages/bishop_white.png")
	ThrowEror(err)
	BishopB, _, _ := ebitenutil.NewImageFromFile("./ChessImages/bishop_black.png")
	ThrowEror(err)
	RookW, _, _ := ebitenutil.NewImageFromFile("./ChessImages/rook_white.png")
	ThrowEror(err)
	RookB, _, _ := ebitenutil.NewImageFromFile("./ChessImages/rook_black.png")
	ThrowEror(err)
	KnightW, _, _ := ebitenutil.NewImageFromFile("./ChessImages/knight_white.png")
	ThrowEror(err)
	KnightB, _, _ := ebitenutil.NewImageFromFile("./ChessImages/knight_black.png")
	ThrowEror(err)
	QueenW, _, _ := ebitenutil.NewImageFromFile("./ChessImages/queen_white.png")
	ThrowEror(err)
	QueenB, _, _ := ebitenutil.NewImageFromFile("./ChessImages/queen_black.png")
	ThrowEror(err)
	KingW, _, _ := ebitenutil.NewImageFromFile("./ChessImages/king_white.png")
	ThrowEror(err)
	KingB, _, _ := ebitenutil.NewImageFromFile("./ChessImages/king_black.png")
	ThrowEror(err)

	// Map image Pieces
	pieceImageMap := map[PieceDescription]*ebiten.Image{
		{PieceName: Pawn, color: "white"}:   PawnW,
		{PieceName: Pawn, color: "black"}:   PawnB,
		{PieceName: Bishop, color: "white"}: BishopW,
		{PieceName: Bishop, color: "black"}: BishopB,
		{PieceName: Rook, color: "white"}:   RookW,
		{PieceName: Rook, color: "black"}:   RookB,
		{PieceName: Knight, color: "white"}: KnightW,
		{PieceName: Knight, color: "black"}: KnightB,
		{PieceName: Queen, color: "white"}:  QueenW,
		{PieceName: Queen, color: "black"}:  QueenB,
		{PieceName: King, color: "white"}:   KingW,
		{PieceName: King, color: "black"}:   KingB,
	}

	return pieceImageMap
}

func ThrowEror(err error) {
	if err != nil {
		log.Fatalf("Couldn't Load image %v", err)
	}
}
