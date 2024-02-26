package gamemain

import "log"

type PieceType int

const (
	Pawn   PieceType = 1
	Bishop           = 2
	Rook             = 3
	Knight           = 4
	Queen            = 5
	King             = 6
)

type PieceDescription struct {
	PieceName PieceType
	color     string
}

func (p PieceType) peices() string {
	switch p {
	case Pawn:
		return "pawn"
	case Bishop:
		return "bishop"
	case Rook:
		return "rook"
	case Knight:
		return "knight"
	case Queen:
		return "queen"
	case King:
		return "king"
	default:
		log.Fatalf("Invalid peice type")
		return ""
	}
}
