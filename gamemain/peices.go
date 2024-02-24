package gamemain

import "github.com/hajimehoshi/ebiten/v2"

/*
Need a struct for each peice for rowPos, colPos, and importanceVal for
future AI built on top of it.
*/

func UpdatePeices() error {
	return nil
}

func DrawPeices(screen *ebiten.Image) {

}

type pawn struct {
	rowPos int
	colPos int
	imVal  int // Value of Importance
}
type king struct {
	rowPos int
	colPos int
	imVal  int // Value of Importance
}
type bishop struct {
	rowPos int
	colPos int
	imVal  int // Value of Importance
}
type knight struct {
	rowPos int
	colPos int
	imVal  int // Value of Importance
}
type rook struct {
	rowPos int
	colPos int
	imVal  int // Value of Importance
}
type queen struct {
	rowPos int
	colPos int
	imVal  int // Value of Importance
}
