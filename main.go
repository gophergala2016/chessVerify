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

	initPawns()
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
	var sourceCol = source[0]
	var sourceRow = source[1]

	//if a white piece is picked and its blacks turn or if a black piece is picked and its whites turn or if no piece is picked up return false

	fmt.Printf("%c %c\n", sourceCol, sourceRow)
	var targetCol = target[0]
	var targetRow = target[1]

	//converting to proper format
	newSourceCol := convertLetter(sourceCol)
	newSourceRow := 8 - int(sourceRow-'0')

	newTargetCol := convertLetter(targetCol)
	newTargetRow := 8 - int(targetRow-'0')

	fmt.Printf("newSourceRow %d newSourceCol %d\n", newSourceRow, newSourceCol)
	fmt.Printf("newTargetRow %d newTargetCol %d\n", newTargetRow, newTargetCol)
	
	
	
	//ensuring a digit is entered into the ChessBoard array to prevent index out of range
	if newSourceRow < 0 || newSourceRow > 7 || newSourceCol < 0 || newSourceCol > 7 || newTargetRow < 0 || newTargetRow > 7 || newTargetCol < 0 || newTargetCol > 7{
		fmt.Println("Invalid input")
		return false
	} 
	
	//identifying the piece that was selected
	piece := ChessBoard[newSourceRow][newSourceCol]
	//piece without color specification
	noColorPiece := fmt.Sprintf("%c", piece[1])
	colorOnly := fmt.Sprintf("%c", piece[0])
	fmt.Println(noColorPiece)

	if (whiteTurn == true && colorOnly == "b") || (whiteTurn == false && colorOnly == "w") || source == "--" {
		return false
	}
	//checking to make sure player doesn't capture his own pieces
	targetSquare := ChessBoard[newTargetRow][newTargetCol]
	targetColor := fmt.Sprintf("%c", targetSquare[0])
	if colorOnly == targetColor{
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
		}

	case "K":
		result := kingMove(newSourceRow, newSourceCol, newTargetRow, newTargetCol)
		if result == false {
			return false
		}
	
	default:
		fmt.Println("Error not valid piece")
		return false

	}

	makeMove(newSourceRow, newSourceCol, newTargetRow, newTargetCol, piece)

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

//makes the chess move on the board, verify the move first
func makeMove(sourceRow int, sourceCol int, targetRow int, targetCol int, piece string) {
	
	//make the source square blank as now the piece is no longer there
	ChessBoard[sourceRow][sourceCol] = "--"
	//if pawn reaches the 8th or 1st rank auto promote to queen for now
	
	if targetRow == 0 && piece == "wP"{
		ChessBoard[targetRow][targetCol] = "wQ"   //white queen promotion
	}else if targetRow == 7 && piece == "bP"{
		ChessBoard[targetRow][targetCol] = "bQ"   //black queen promtion
	}else{
		ChessBoard[targetRow][targetCol] = piece  //place the piece to its new target square
	}
	 
	switchTurns()
	
}

func switchTurns(){
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