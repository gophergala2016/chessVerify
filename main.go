// Jonathan Chin
// 1/24/2016
// Gopher Gala 2016
// verify if chess move is legal by using source and target notation Ex: e2 e4
package main

import (
	"fmt"
)

//keeps track of whose move it is, true means its whites turn and false means its blacks turn
var whiteTurn = true

var ChessBoard = [][]string{
	[]string{"bR", "bN", "bB", "bQ", "bK", "bB", "bN", "bR"},
	[]string{"bP", "bP", "bP", "bP", "bP", "bP", "bP", "bP"},
	[]string{"--", "--", "--", "--", "--", "--", "--", "--"},
	[]string{"--", "--", "--", "--", "--", "--", "--", "--"},
	[]string{"--", "--", "--", "--", "--", "--", "--", "--"},
	[]string{"--", "--", "--", "--", "--", "--", "--", "--"},
	[]string{"wP", "wP", "wP", "wP", "wP", "wP", "wP", "wP"},
	[]string{"wR", "wN", "wB", "wQ", "wK", "wB", "wN", "wR"},
}

//stores array of booleans indicating whether or not the pawns have made their first move yet
var whitePawns [8]bool
var blackPawns [8]bool

//stores array to indicate whether or not the pawns can perform an enpassent
var whitePass [8]bool
var blackPass [8]bool

func main() {

	initGame()
	printBoard()
	var result bool
	var source string
	var target string
	fmt.Println("Welcome to chessVerify, please input your chess move and press enter")
	for {
		fmt.Scanf("%s %s\n", &source, &target)
		result = chessVerify(source, target)

		if result == true {
			fmt.Println("Valid move")
			printBoard()

		} else {
			fmt.Println("Invalid move")
		}
	}

}

//returns true if the move is valid otherwise it returns false
func chessVerify(source string, target string) bool {
	if len(source) != 2 {
		fmt.Println("Invalid input length")
		return false
	}
	var sourceCol = source[0]
	var sourceRow = source[1]

	//if a white piece is picked and its blacks turn or if a black piece is picked and its whites turn or if no piece is picked up return false

	//checking length to ensure no index out of range error
	if len(target) != 2 {
		fmt.Println("Invalid input length")
		return false
	}
	var targetCol = target[0]
	var targetRow = target[1]

	//converting to proper format
	newSourceCol := convertLetter(sourceCol)
	newSourceRow := 8 - int(sourceRow-'0')

	newTargetCol := convertLetter(targetCol)
	newTargetRow := 8 - int(targetRow-'0')

	//ensuring a digit is entered into the ChessBoard array to prevent index out of range
	if newSourceRow < 0 || newSourceRow > 7 || newSourceCol < 0 || newSourceCol > 7 || newTargetRow < 0 || newTargetRow > 7 || newTargetCol < 0 || newTargetCol > 7 {
		fmt.Println("Invalid input")
		return false
	}

	//identifying the piece that was selected
	piece := ChessBoard[newSourceRow][newSourceCol]
	//piece without color specification
	noColorPiece := piece[1:2]
	colorOnly := piece[0:1]

	if (whiteTurn == true && colorOnly == "b") || (whiteTurn == false && colorOnly == "w") || source == "--" {
		return false
	}
	//checking to make sure player doesn't capture his own pieces
	targetSquare := ChessBoard[newTargetRow][newTargetCol]
	targetColor := targetSquare[0:1]

	if colorOnly == targetColor {
		fmt.Println("You can't capture your own piece.")
		return false
	}

	//verifying the piece move
	switch noColorPiece {
	case "P":
		if piece == "wP" {
			result := whitePawnMove(newSourceRow, newSourceCol, newTargetRow, newTargetCol)
			if result == false {
				return false
			}

		} else {

			result := blackPawnMove(newSourceRow, newSourceCol, newTargetRow, newTargetCol)
			if result == false {
				return false
			}
		}

	case "N":
		result := knightMove(newSourceRow, newSourceCol, newTargetRow, newTargetCol)
		if result == false {
			return false
		}

	case "B":
		result := bishopMove(newSourceRow, newSourceCol, newTargetRow, newTargetCol)
		if result == false {
			return false
		}

	case "Q":
		result := queenMove(newSourceRow, newSourceCol, newTargetRow, newTargetCol)
		if result == false {
			return false
		}

	case "R":
		result := rookMove(newSourceRow, newSourceCol, newTargetRow, newTargetCol)
		if result == false {
			return false
		}else{
			rookUpdate = true  //used to indicate if a rook has moved, used for castling rights
		}

	case "K":
		result := kingMove(newSourceRow, newSourceCol, newTargetRow, newTargetCol)
		if result == false {
			return false
		} else { //if its valid king move then update coordinates in the global variables which keeps track of kings location
			kingUpdate = true
		}

	default:
		fmt.Println("Error not valid piece")
		return false

	}

	capturedPiece := makeMove(newSourceRow, newSourceCol, newTargetRow, newTargetCol, piece)
	
	//if the player made a move and his king can be captured that move has to be undone and return false as he didn't stop the check
	if whiteTurn == true && isWhiteInCheck() == true {
		undoMove(newSourceRow, newSourceCol, newTargetRow, newTargetCol, piece, capturedPiece)
		fmt.Println("White cannot make that move as they are in check")
		return false
	} else if whiteTurn == false && isBlackInCheck() == true {
		undoMove(newSourceRow, newSourceCol, newTargetRow, newTargetCol, piece, capturedPiece)
		fmt.Println("Black cannot make that move as they are in check")
		return false
	}
	if kingUpdate == true { //updating new location of king for quick access
		if colorOnly == "b" {
			blackKingX = newTargetRow
			blackKingY = newTargetCol
			bkMoved = true  //can no longer castle with black king
		} else if colorOnly == "w" {
			whiteKingX = newTargetRow
			whiteKingY = newTargetCol
			wkMoved = true   //can no longer castle with white king
		} else {
			fmt.Println("Invalid king color")
		}
		kingUpdate = false
	}
	
	if rookUpdate == true{
		if piece == "bR" && newSourceRow == 0 && newSourceCol== 0{ //black queen rook
			bqrMoved = true
		}else if piece == "bR" && newSourceRow == 0 && newSourceCol== 7{ //black king rook
			bkrMoved = true
		}else if piece == "wR" && newSourceRow == 7 && newSourceCol== 0{ //white queen rook
			wqrMoved = true
		}else if piece == "wR" && newSourceRow == 7 && newSourceCol== 7{ //white king rook move
			wkrMoved = true
		}
		rookUpdate = false
	}
	switchTurns()
	return true
}

//changes chess letter notation to a number a=0, b=1, c=2, etc
func convertLetter(letter byte) int {
	switch letter {
	case 'a': //this is a file on chess board
		return 0
	case 'b':
		return 1
	case 'c':
		return 2
	case 'd':
		return 3
	case 'e':
		return 4
	case 'f':
		return 5
	case 'g':
		return 6
	case 'h':
		return 7
	default:
		fmt.Println("Invalid file on chess board")
		return -1
	}

}

//makes the chess move on the board, verify the move first, returns captured piece as a string to be used in case of a move undo
func makeMove(sourceRow int, sourceCol int, targetRow int, targetCol int, piece string) string {

	capturedPiece := ChessBoard[sourceRow][sourceCol]
	//make the source square blank as now the piece is no longer there
	ChessBoard[sourceRow][sourceCol] = "--"
	//if pawn reaches the 8th or 1st rank auto promote to queen for now

	if targetRow == 0 && piece == "wP" {
		ChessBoard[targetRow][targetCol] = "wQ" //white queen promotion
	} else if targetRow == 7 && piece == "bP" {
		ChessBoard[targetRow][targetCol] = "bQ" //black queen promtion
	} else {
		ChessBoard[targetRow][targetCol] = piece //place the piece to its new target square
	}
	return capturedPiece

}

func switchTurns() {
	if whiteTurn == true {
		whiteTurn = false
	} else {
		whiteTurn = true
	}
}

func printBoard() {

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Printf("%s ", ChessBoard[i][j])
		}
		fmt.Printf("\n")
	}
}

//intitalize all pawns to false as they have not moved yet, and also initialize all en passent to false
func initGame() {
	for i := 0; i < 8; i++ {
		whitePawns[i] = false
		blackPawns[i] = false
		whitePass[i] = false
		blackPass[i] = false
	}
	//castling rights init for kings
	wkMoved = false
	bkMoved = false
	//castling rights int for rooks
	wkrMoved = false
	wqrMoved = false
	bkrMoved = false
	bqrMoved = false
	//storing coordinates for starting location of both kings, X is row and Y is col
	whiteKingX = 7
	whiteKingY = 4
	blackKingX = 0
	blackKingY = 4

	kingUpdate = false
	rookUpdate = false
}

//undo a move if a player makes a move and doesn't stop check
func undoMove(sourceRow int, sourceCol int, targetRow int, targetCol int, piece string, capturedPiece string) {
	ChessBoard[sourceRow][sourceCol] = piece
	ChessBoard[targetRow][targetCol] = capturedPiece
}
