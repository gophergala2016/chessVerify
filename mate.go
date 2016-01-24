package main

import (
	"fmt"
)

//runs through the entire ChessBoard array and searches for black pieces and brute all their possible moves
//to see if it can capture the white king in one move
func isWhiteInCheck() bool {
	result := canBlackKillSquare(whiteKingX, whiteKingY)
	if result == true { //then white's king is in check
		return true
	}
	return false
}

func isBlackInCheck() bool {
	result := canWhiteKillSquare(blackKingX, blackKingY)
	if result == true { //then black's king is in check
		return true
	}
	return false
}

//if white has no legal moves its stalemate
func isWhiteStaleMate() {

}

func isBlackStaleMate() {

}

//if white's king is in check and he has no squares to run too
func isWhiteMated() {

}

func isBlackMated() {

}

//checks if a square is attacked by white in one turn, used to identify check and checkmates
func canWhiteKillSquare(targetRow int, targetCol int) bool {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			color := ChessBoard[i][j]
			if color[0:1] == "w" {
				switch color[1:2] {
				case "P":
					result := whitePawnAttack(i, j, targetRow, targetCol)
					if result == true {
						return true
					}
				case "N":
					result := knightAttack(i, j, targetRow, targetCol)
					if result == true {
						return true
					}
				case "B":
					result := bishopAttack(i, j, targetRow, targetCol)
					if result == true {
						return true
					}
				case "R":
					result := rookAttack(i, j, targetRow, targetCol)
					if result == true {
						return true
					}
				case "Q":
					result := queenAttack(i, j, targetRow, targetCol)
					if result == true {
						return true
					}
				case "K":
					result := kingAttack(i, j, targetRow, targetCol)
					if result == true {
						return true
					}

				default:
					fmt.Println("Invalid piece type")
				}
			}
		}
	}
	return false
}

func canBlackKillSquare(targetRow int, targetCol int) bool {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			color := ChessBoard[i][j]
			if color[0:1] == "b" {
				switch color[1:2] {
				case "P":
					result := blackPawnAttack(i, j, targetRow, targetCol)
					if result == true {
						return true
					}
				case "N":
					result := knightAttack(i, j, targetRow, targetCol)
					if result == true {
						return true
					}
				case "B":
					result := bishopAttack(i, j, targetRow, targetCol)
					if result == true {
						return true
					}
				case "R":
					result := rookAttack(i, j, targetRow, targetCol)
					if result == true {
						return true
					}
				case "Q":
					result := queenAttack(i, j, targetRow, targetCol)
					if result == true {
						return true
					}
				case "K":
					result := kingAttack(i, j, targetRow, targetCol)
					if result == true {
						return true
					}
				default:
					fmt.Println("Invalid piece type")
				}
			}
		}
	}
	return false
}
