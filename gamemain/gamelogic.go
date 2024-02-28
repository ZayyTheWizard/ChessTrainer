package gamemain

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func (inf *BoardDimension) highlightPiece(screen *ebiten.Image, row, col int) {
	screenWidth, screenHeight := screen.Bounds().Dx(), screen.Bounds().Dy()

	squareSize := int(math.Min(float64(screenWidth)/float64(inf.COL), float64(screenHeight)/float64(inf.ROW)))
	highlightColor := color.RGBA{80, 80, 0, 0} // Red color for highlighting
	highlightSquare := ebiten.NewImage(squareSize, squareSize)
	highlightSquare.Fill(highlightColor)

	//var x, y = ebiten.CursorPosition()
	if (inf.isFieldNotEmpty(posMap[RowCol{row: row, col: col}])) && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		// Highlight Piece
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(col*squareSize), float64(row*squareSize))
		screen.DrawImage(highlightSquare, opts)
	}
}

func (inf *BoardDimension) isFieldNotEmpty(pos int) bool {
	return inf.consoleRepresentation[pos].PieceName != 0
}

func (inf *BoardDimension) isFieldEmpty(pos int) bool {
	return inf.consoleRepresentation[pos].PieceName == 0
}

func (inf *BoardDimension) mouseLoc(screen *ebiten.Image) (int, int) {
	screenWidth, screenHeight := screen.Bounds().Dx(), screen.Bounds().Dy()

	x, y := ebiten.CursorPosition()
	squareSize := int(math.Min(float64(screenWidth)/float64(inf.COL), float64(screenHeight)/float64(inf.ROW)))
	row := y / squareSize
	col := x / squareSize
	return row, col
}

var (
	posMap = map[RowCol]int{
		{row: 0, col: 0}: 0,
		{row: 0, col: 1}: 1,
		{row: 0, col: 2}: 2,
		{row: 0, col: 3}: 3,
		{row: 0, col: 4}: 4,
		{row: 0, col: 5}: 5,
		{row: 0, col: 6}: 6,
		{row: 0, col: 7}: 7,
		{row: 1, col: 0}: 8,
		{row: 1, col: 1}: 9,
		{row: 1, col: 2}: 10,
		{row: 1, col: 3}: 11,
		{row: 1, col: 4}: 12,
		{row: 1, col: 5}: 13,
		{row: 1, col: 6}: 14,
		{row: 1, col: 7}: 15,
		{row: 2, col: 0}: 16,
		{row: 2, col: 1}: 17,
		{row: 2, col: 2}: 18,
		{row: 2, col: 3}: 19,
		{row: 2, col: 4}: 20,
		{row: 2, col: 5}: 21,
		{row: 2, col: 6}: 22,
		{row: 2, col: 7}: 23,
		{row: 3, col: 0}: 24,
		{row: 3, col: 1}: 25,
		{row: 3, col: 2}: 26,
		{row: 3, col: 3}: 27,
		{row: 3, col: 4}: 28,
		{row: 3, col: 5}: 29,
		{row: 3, col: 6}: 30,
		{row: 3, col: 7}: 31,
		{row: 4, col: 0}: 32,
		{row: 4, col: 1}: 33,
		{row: 4, col: 2}: 34,
		{row: 4, col: 3}: 35,
		{row: 4, col: 4}: 36,
		{row: 4, col: 5}: 37,
		{row: 4, col: 6}: 38,
		{row: 4, col: 7}: 39,
		{row: 5, col: 0}: 40,
		{row: 5, col: 1}: 41,
		{row: 5, col: 2}: 42,
		{row: 5, col: 3}: 43,
		{row: 5, col: 4}: 44,
		{row: 5, col: 5}: 45,
		{row: 5, col: 6}: 46,
		{row: 5, col: 7}: 47,
		{row: 6, col: 0}: 48,
		{row: 6, col: 1}: 49,
		{row: 6, col: 2}: 50,
		{row: 6, col: 3}: 51,
		{row: 6, col: 4}: 52,
		{row: 6, col: 5}: 53,
		{row: 6, col: 6}: 54,
		{row: 6, col: 7}: 55,
		{row: 7, col: 0}: 56,
		{row: 7, col: 1}: 57,
		{row: 7, col: 2}: 58,
		{row: 7, col: 3}: 59,
		{row: 7, col: 4}: 60,
		{row: 7, col: 5}: 61,
		{row: 7, col: 6}: 62,
		{row: 7, col: 7}: 63,
	}
)

type RowCol struct {
	row, col int
}

func (inf *BoardDimension) gameLogic(screen *ebiten.Image, s *selectedPiece) {
	inf.movePawn(screen, s)
}

func (inf *BoardDimension) movePawn(screen *ebiten.Image, s *selectedPiece) {
	var row, col = inf.mouseLoc(screen)
	/*
		1. Check who move it is White or black
		2. Check if king is in check
	*/
	if s.row == 6 || s.row == 1 { // If it is the Pawn's first move it can move the 1 or 2 squares foward
		if s.color == "white" && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) { // White Pawns
			if row == s.row-2 && col == s.col {
				inf.consoleRepresentation[posMap[RowCol{row: s.row, col: s.col}]-16] = inf.consoleRepresentation[posMap[RowCol{row: s.row, col: s.col}]]
				inf.consoleRepresentation[posMap[RowCol{row: s.row, col: s.col}]] = PieceDescription{}
				s = &selectedPiece{Piece: 0, row: 0, col: 0, color: ""}
			} else if row == s.row-1 && col == s.col {
				inf.consoleRepresentation[posMap[RowCol{row: s.row, col: s.col}]-8] = inf.consoleRepresentation[posMap[RowCol{row: s.row, col: s.col}]]
				inf.consoleRepresentation[posMap[RowCol{row: s.row, col: s.col}]] = PieceDescription{}
				s = &selectedPiece{Piece: 0, row: 0, col: 0, color: ""}
			}

		}
		if s.color == "black" && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) { // Black Pawns
			if row == s.row+2 && col == s.col {
				inf.consoleRepresentation[posMap[RowCol{row: s.row, col: s.col}]+16] = inf.consoleRepresentation[posMap[RowCol{row: s.row, col: s.col}]]
				inf.consoleRepresentation[posMap[RowCol{row: s.row, col: s.col}]] = PieceDescription{}
				s = &selectedPiece{Piece: 0, row: 0, col: 0, color: ""}
			} else if row == s.row+1 && col == s.col {
				inf.consoleRepresentation[posMap[RowCol{row: s.row, col: s.col}]+8] = inf.consoleRepresentation[posMap[RowCol{row: s.row, col: s.col}]]
				inf.consoleRepresentation[posMap[RowCol{row: s.row, col: s.col}]] = PieceDescription{}
				s = &selectedPiece{Piece: 0, row: 0, col: 0, color: ""}
			}
		}
	}
}

func (inf *BoardDimension) highlightPawnMove(screen ebiten.Image) {
	/*
		1. Check who move it is White or black
		2. Check if king is in check
	*/
}

type selectedPiece struct {
	Piece    PieceType
	row, col int
	color    string
}

func (inf BoardDimension) updateCurrentPiece(row, col int, s *selectedPiece) {
	if (inf.isFieldNotEmpty(posMap[RowCol{row: row, col: col}])) && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		// Highlight Piece
		s.Piece = inf.consoleRepresentation[posMap[RowCol{row: row, col: col}]].PieceName
		s.color = inf.consoleRepresentation[posMap[RowCol{row: row, col: col}]].color
		s.row = row
		s.col = col
	}
}
