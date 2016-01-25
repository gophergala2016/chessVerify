package main

import (
	"fmt"
)

//if king moved or not
var wkMoved bool
var bkMoved bool

var wkrMoved bool
var wqrMoved bool
var bkrMoved bool
var bqrMoved bool

//stores location of king for easy access
var whiteKingX int
var whiteKingY int
var blackKingX int
var blackKingY int

//used to figure out if king position needs to be updated
var kingUpdate bool
var rookUpdate bool //castling rights

func knightMove(sourceRow int, sourceCol int, targetRow int, targetCol int) bool {
	if (targetRow-sourceRow == 2 || targetRow-sourceRow == -2) && (targetCol-sourceCol == 1 || targetCol-sourceCol == -1) {
		return true
	} else if (targetRow-sourceRow == 1 || targetRow-sourceRow == -1) && (targetCol-sourceCol == 2 || targetCol-sourceCol == -2) {
		return true
	}
	return false
}

func bishopMove(sourceRow int, sourceCol int, targetRow int, targetCol int) bool {

	if targetRow < sourceRow && targetCol < sourceCol && sourceRow-targetRow == sourceCol-targetCol { //left up
		for i, j := sourceRow-1, sourceCol-1; i > targetRow; i, j = i-1, j-1 {
			if ChessBoard[i][j] != "--" {
				fmt.Println("Piece is blocking the bishop")
				return false
			}
		}

	} else if targetRow > sourceRow && targetCol < sourceCol && targetRow-sourceRow == sourceCol-targetCol { //left down
		for i, j := sourceRow+1, sourceCol-1; i < targetRow; i, j = i+1, j-1 {
			if ChessBoard[i][j] != "--" {
				fmt.Println("Piece is blocking the bishop")
				return false
			}
		}
	} else if targetRow > sourceRow && targetCol > sourceCol && targetRow-sourceRow == targetCol-sourceCol { //right down
		for i, j := sourceRow+1, sourceCol+1; i < targetRow; i, j = i+1, j+1 {
			if ChessBoard[i][j] != "--" {
				fmt.Println("Piece is blocking the bishop")
				return false
			}
		}
	} else if targetRow < sourceRow && targetCol > sourceCol && sourceRow-targetRow == targetCol-sourceCol { //right up
		for i, j := sourceRow-1, sourceCol+1; i < targetRow; i, j = i-1, j+1 {
			if ChessBoard[i][j] != "--" {
				fmt.Println("Piece is blocking the bishop")
				return false
			}
		}
	} else {
		fmt.Println("Invalid bishop move")
		return false
	}
	return true
}

//combination of bishop and rook
func queenMove(sourceRow int, sourceCol int, targetRow int, targetCol int) bool {

	//check diagonals using bishop moves
	if targetRow < sourceRow && targetCol < sourceCol && sourceRow-targetRow == sourceCol-targetCol { //left up
		for i, j := sourceRow-1, sourceCol-1; i > targetRow; i, j = i-1, j-1 {
			if ChessBoard[i][j] != "--" {
				fmt.Println("Piece is blocking the queen")
				return false
			}
		}

	} else if targetRow > sourceRow && targetCol < sourceCol && targetRow-sourceRow == sourceCol-targetCol { //left down
		for i, j := sourceRow+1, sourceCol-1; i < targetRow; i, j = i+1, j-1 {
			if ChessBoard[i][j] != "--" {
				fmt.Println("Piece is blocking the queen")
				return false
			}
		}
	} else if targetRow > sourceRow && targetCol > sourceCol && targetRow-sourceRow == targetCol-sourceCol { //right down
		for i, j := sourceRow+1, sourceCol+1; i < targetRow; i, j = i+1, j+1 {
			if ChessBoard[i][j] != "--" {
				fmt.Println("Piece is blocking the queen")
				return false
			}
		}
	} else if targetRow < sourceRow && targetCol > sourceCol && sourceRow-targetRow == targetCol-sourceCol { //right up
		for i, j := sourceRow-1, sourceCol+1; i < targetRow; i, j = i-1, j+1 {
			if ChessBoard[i][j] != "--" {
				fmt.Println("Piece is blocking the queen")
				return false
			}
		}
	} else if sourceRow < targetRow && sourceCol == targetCol { //up
		for i := sourceRow + 1; i < targetRow; i++ {
			if ChessBoard[i][sourceCol] != "--" {
				fmt.Println("Piece is blocking the queen")
				return false
			}
		}

	} else if sourceRow > targetRow && sourceCol == targetCol { //down
		for i := sourceRow - 1; i > targetRow; i-- {
			if ChessBoard[i][sourceCol] != "--" {
				fmt.Println("Piece is blocking the queen")
				return false
			}
		}
	} else if targetCol > sourceCol && sourceRow == targetRow { //right
		for i := sourceCol + 1; i < targetCol; i++ {
			if ChessBoard[sourceRow][i] != "--" {
				fmt.Println("Piece is blocking the queen")
				return false
			}
		}
	} else if targetCol < sourceCol && sourceRow == targetRow { //left
		for i := sourceCol - 1; i > targetCol; i-- {
			if ChessBoard[sourceRow][i] != "--" {
				fmt.Println("Piece is blocking the queen")
				return false
			}
		}
	} else {
		fmt.Println("Invalid queen move 2")
		return false
	}
	return true
}

func rookMove(sourceRow int, sourceCol int, targetRow int, targetCol int) bool {
	if sourceRow < targetRow && sourceCol == targetCol { //up
		for i := sourceRow + 1; i < targetRow; i++ {
			if ChessBoard[i][sourceCol] != "--" {
				fmt.Println("Piece is blocking the rook")
				return false
			}
		}

	} else if sourceRow > targetRow && sourceCol == targetCol { //down
		for i := sourceRow - 1; i > targetRow; i-- {
			if ChessBoard[i][sourceCol] != "--" {
				fmt.Println("Piece is blocking the rook")
				return false
			}
		}
	} else if targetCol > sourceCol && sourceRow == targetRow { //right
		for i := sourceCol + 1; i < targetCol; i++ {
			if ChessBoard[sourceRow][i] != "--" {
				fmt.Println("Piece is blocking the rook")
				return false
			}
		}
	} else if targetCol < sourceCol && sourceRow == targetRow { //left
		for i := sourceCol - 1; i > targetCol; i-- {
			if ChessBoard[sourceRow][i] != "--" {
				fmt.Println("Piece is blocking the rook")
				return false
			}
		}
	} else {
		fmt.Println("Invalid rook move")
		return false
	}

	return true
}

//if king moves two spaces to right as white it checks for kingside castle, three spaces to the left as white checks for white queen castle
func kingMove(sourceRow int, sourceCol int, targetRow int, targetCol int) bool {
	if sourceRow == targetRow && (sourceCol-targetCol == 1 || sourceCol-targetCol == -1) { //left or right
		return true //make sure king doesn't walk into check, have a function which checks a color if it can attack that square
	} else if sourceCol == targetCol && (sourceRow-targetRow == 1 || sourceRow-targetRow == -1) { //up or down
		return true
	} else if (sourceCol-targetCol == 1 || sourceCol-targetCol == -1) && (sourceRow-targetRow == 1 || sourceRow-targetRow == -1) { //diagonals
		return true
	} else if ChessBoard[sourceRow][sourceCol] == "wK" && ChessBoard[7][6] == "--" && ChessBoard[7][5] == "--" && sourceRow == targetRow && targetCol-sourceCol == 2 && canWhiteCastleKing() == true { //white king castle
		fmt.Println("White kingside castle")
		ChessBoard[7][7] = "--" //removing rook off of h file
		ChessBoard[7][5] = "wR" //moving rook as well
		return true
	} else if ChessBoard[sourceRow][sourceCol] == "wK" && ChessBoard[7][1] == "--" && ChessBoard[7][2] == "--" && ChessBoard[7][3] == "--" && sourceRow == targetRow && sourceCol-targetCol == 2 && canWhiteCastleQueen() == true { //white queen castle
		fmt.Println("White queenside castle")
		ChessBoard[7][0] = "--" //removing rook off of a file
		ChessBoard[7][3] = "wR" //moving rook as well
		return true
	} else if ChessBoard[sourceRow][sourceCol] == "bK" && ChessBoard[0][6] == "--" && ChessBoard[0][5] == "--" && sourceRow == targetRow && targetCol-sourceCol == 2 && canBlackCastleKing() == true { //black king castle
		fmt.Println("Black kingside castle")
		ChessBoard[0][7] = "--" //removing rook off of h file
		ChessBoard[0][5] = "bR" //moving rook as well
		return true
	} else if ChessBoard[sourceRow][sourceCol] == "bK" && ChessBoard[0][1] == "--" && ChessBoard[0][2] == "--" && ChessBoard[0][3] == "--" && sourceRow == targetRow && sourceCol-targetCol == 2 && canBlackCastleQueen() == true { //black queenside castle
		fmt.Println("Black queenside castle")
		ChessBoard[0][0] = "--" //removing rook off of a file
		ChessBoard[0][3] = "bR" //moving rook as well
		return true
	}

	fmt.Println("Invalid king move")
	return false
}
