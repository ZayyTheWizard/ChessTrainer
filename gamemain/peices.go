package gamemain

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type PieceType int

const (
	Pawn   PieceType = 1
	Bishop PieceType = 2
	Rook   PieceType = 3
	Knight PieceType = 4
	Queen  PieceType = 5
	King   PieceType = 6
)

type PieceDescription struct {
	PieceName PieceType
	color     string
}

func (p PieceType) PeiceNameAsString() string {
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

// load piece images here
func initImages() map[PieceDescription]*ebiten.Image {
	PawnW, _, _ := ebitenutil.NewImageFromFile("./ChessImages/pawn_white.png")
	PawnB, _, _ := ebitenutil.NewImageFromFile("./ChessImages/pawn_black.png")
	BishopW, _, _ := ebitenutil.NewImageFromFile("./ChessImages/bishop_white.png")
	BishopB, _, _ := ebitenutil.NewImageFromFile("./ChessImages/bishop_black.png")
	RookW, _, _ := ebitenutil.NewImageFromFile("./ChessImages/rook_white.png")
	RookB, _, _ := ebitenutil.NewImageFromFile("./ChessImages/rook_black.png")
	KnightW, _, _ := ebitenutil.NewImageFromFile("./ChessImages/knight_white.png")
	KnightB, _, _ := ebitenutil.NewImageFromFile("./ChessImages/knight_black.png")
	QueenW, _, _ := ebitenutil.NewImageFromFile("./ChessImages/queen_white.png")
	QueenB, _, _ := ebitenutil.NewImageFromFile("./ChessImages/queen_black.png")
	KingW, _, _ := ebitenutil.NewImageFromFile("./ChessImages/king_white.png")
	KingB, _, _ := ebitenutil.NewImageFromFile("./ChessImages/king_black.png")

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
