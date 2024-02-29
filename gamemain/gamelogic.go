package gamemain

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func (inf *BoardDimension) isFieldNotEmpty(pos int) bool {
	return inf.consoleRepresentation[pos].PieceName != 0
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

func (inf *BoardDimension) gameLogic(screen *ebiten.Image, s *selectedPiece, turn *string) {
	inf.movePawn(screen, s, turn)
}

// Pawn Logic Start
func (inf *BoardDimension) promotePawn() {

	for col := 0; col < 8; col++ {
		if inf.consoleRepresentation[posMap[RowCol{row: 0, col: col}]].PieceName == Pawn && inf.consoleRepresentation[posMap[RowCol{row: 0, col: col}]].color == "white" {
			inf.consoleRepresentation[posMap[RowCol{row: 0, col: col}]] = PieceDescription{PieceName: Queen, color: "white"}
		} else if inf.consoleRepresentation[posMap[RowCol{row: 7, col: col}]].PieceName == Pawn && inf.consoleRepresentation[posMap[RowCol{row: 0, col: col}]].color == "black" {
			inf.consoleRepresentation[posMap[RowCol{row: 7, col: col}]] = PieceDescription{PieceName: Queen, color: "black"}
		}
	}
}

// Generate moves array for any given turn given

func (inf *BoardDimension) movePawn(screen *ebiten.Image, s *selectedPiece, turn *string) {
	var mouseRow, mouseCol = inf.mouseLoc(screen)
	inf.promotePawn()
	/*
		Check if king is in check {
			Can move block check or capture checker {
				handle that
			} else {
				return;
			}
		}
	*/
	if *turn == "white" {
		if s.row == 6 { // If it is the Pawn's first move it can move the 1 or 2 squares foward
			if s.color == "white" && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) { // White Pawns
				if mouseRow == s.row-2 && mouseCol == s.col && inf.consoleRepresentation[posMap[RowCol{row: s.row - 2, col: s.col}]].PieceName == 0 {
					inf.consoleRepresentation[posMap[RowCol{row: s.row - 2, col: s.col}]] = inf.consoleRepresentation[posMap[RowCol{row: s.row, col: s.col}]]
					inf.consoleRepresentation[posMap[RowCol{row: s.row, col: s.col}]] = PieceDescription{}
					s = &selectedPiece{Piece: 0, row: 0, col: 0, color: ""}
					*turn = "black"
				} else if mouseRow == s.row-1 && mouseCol == s.col && inf.consoleRepresentation[posMap[RowCol{row: s.row - 1, col: s.col}]].PieceName == 0 {
					inf.consoleRepresentation[posMap[RowCol{row: s.row - 1, col: s.col}]] = inf.consoleRepresentation[posMap[RowCol{row: s.row, col: s.col}]]
					inf.consoleRepresentation[posMap[RowCol{row: s.row, col: s.col}]] = PieceDescription{}
					s = &selectedPiece{Piece: 0, row: 0, col: 0, color: ""}
					*turn = "black"
				}
			}
		} else if s.row != 6 {
			if s.color == "white" && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) { // White Pawns
				if mouseRow == s.row-1 && mouseCol == s.col && inf.consoleRepresentation[posMap[RowCol{row: s.row - 1, col: s.col}]].PieceName == 0 {
					inf.consoleRepresentation[posMap[RowCol{row: s.row - 1, col: s.col}]] = inf.consoleRepresentation[posMap[RowCol{row: s.row, col: s.col}]]
					inf.consoleRepresentation[posMap[RowCol{row: s.row, col: s.col}]] = PieceDescription{}
					s = &selectedPiece{Piece: 0, row: 0, col: 0, color: ""}
					*turn = "black"
				} else if mouseRow == s.row-1 && mouseCol == s.col-1 && inf.consoleRepresentation[posMap[RowCol{row: s.row - 1, col: s.col - 1}]].color == "black" { // white attack on black
					inf.consoleRepresentation[posMap[RowCol{row: s.row - 1, col: s.col - 1}]] = inf.consoleRepresentation[posMap[RowCol{row: s.row, col: s.col}]]
					inf.consoleRepresentation[posMap[RowCol{row: s.row, col: s.col}]] = PieceDescription{}
					s = &selectedPiece{Piece: 0, row: 0, col: 0, color: ""}
					*turn = "black"
				} else if mouseRow == s.row-1 && mouseCol == s.col+1 && inf.consoleRepresentation[posMap[RowCol{row: s.row - 1, col: s.col + 1}]].color == "black" { // white attack on black
					inf.consoleRepresentation[posMap[RowCol{row: s.row - 1, col: s.col + 1}]] = inf.consoleRepresentation[posMap[RowCol{row: s.row, col: s.col}]]
					inf.consoleRepresentation[posMap[RowCol{row: s.row, col: s.col}]] = PieceDescription{}
					s = &selectedPiece{Piece: 0, row: 0, col: 0, color: ""}
					*turn = "black"
				}
			}
		}

	} else if *turn == "black" {
		if s.row == 1 { // If it is the Pawn's first move it can move the 1 or 2 squares foward
			if s.color == "black" && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) { // White Pawns
				if mouseRow == s.row+2 && mouseCol == s.col && inf.consoleRepresentation[posMap[RowCol{row: s.row + 2, col: s.col}]].PieceName == 0 {
					inf.consoleRepresentation[posMap[RowCol{row: s.row + 2, col: s.col}]] = inf.consoleRepresentation[posMap[RowCol{row: s.row, col: s.col}]]
					inf.consoleRepresentation[posMap[RowCol{row: s.row, col: s.col}]] = PieceDescription{}
					s = &selectedPiece{Piece: 0, row: 0, col: 0, color: ""}
					*turn = "white"
				} else if mouseRow == s.row+1 && mouseCol == s.col && inf.consoleRepresentation[posMap[RowCol{row: s.row + 1, col: s.col}]].PieceName == 0 {
					inf.consoleRepresentation[posMap[RowCol{row: s.row + 1, col: s.col}]] = inf.consoleRepresentation[posMap[RowCol{row: s.row, col: s.col}]]
					inf.consoleRepresentation[posMap[RowCol{row: s.row, col: s.col}]] = PieceDescription{}
					s = &selectedPiece{Piece: 0, row: 0, col: 0, color: ""}
					*turn = "white"
				}
			}
		} else if s.row != 1 {
			if s.color == "black" && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) { // White Pawns
				if mouseRow == s.row+1 && mouseCol == s.col && inf.consoleRepresentation[posMap[RowCol{row: mouseRow, col: mouseCol}]].PieceName == 0 {
					inf.consoleRepresentation[posMap[RowCol{row: s.row + 1, col: s.col}]] = inf.consoleRepresentation[posMap[RowCol{row: s.row, col: s.col}]]
					inf.consoleRepresentation[posMap[RowCol{row: s.row, col: s.col}]] = PieceDescription{}
					s = &selectedPiece{Piece: 0, row: 0, col: 0, color: ""}
					*turn = "white"
				} else if mouseRow == s.row+1 && mouseCol == s.col-1 && inf.consoleRepresentation[posMap[RowCol{row: s.row + 1, col: s.col - 1}]].color == "white" { // white attack on black
					inf.consoleRepresentation[posMap[RowCol{row: s.row + 1, col: s.col - 1}]] = inf.consoleRepresentation[posMap[RowCol{row: s.row, col: s.col}]]
					inf.consoleRepresentation[posMap[RowCol{row: s.row, col: s.col}]] = PieceDescription{}
					s = &selectedPiece{Piece: 0, row: 0, col: 0, color: ""}
					*turn = "white"
				} else if mouseRow == s.row+1 && mouseCol == s.col+1 && inf.consoleRepresentation[posMap[RowCol{row: s.row + 1, col: s.col + 1}]].color == "white" { // white attack on black
					inf.consoleRepresentation[posMap[RowCol{row: s.row + 1, col: s.col + 1}]] = inf.consoleRepresentation[posMap[RowCol{row: s.row, col: s.col}]]
					inf.consoleRepresentation[posMap[RowCol{row: s.row, col: s.col}]] = PieceDescription{}
					s = &selectedPiece{Piece: 0, row: 0, col: 0, color: ""}
					*turn = "white"
				}
			}
		}
	}
}

// Pawn Logic End

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
