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

	//identifying the piece that was selected
	piece := ChessBoard[newSourceRow][newSourceCol]
	//piece without color specification
	noColorPiece := fmt.Sprintf("%c", piece[1])
	colorOnly := fmt.Sprintf("%c", piece[0])
	fmt.Println(noColorPiece)

	if (whiteTurn == true && colorOnly == "b") || (whiteTurn == false && colorOnly == "w") || source == "--" {
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
		break
	case "N":
		break
	case "B":
		break
	case "Q":
		break
	case "R":
		break
	case "K":
		break
	default:
		fmt.Println("Error not valid piece")
		return false

	}

	makeMove(newSourceRow, newSourceCol, newTargetRow, newTargetCol, piece)

	return true
}

//intitalize all pawns to false as they have not moved yet, and also initialize all en passent to false
func initPawns() {
	for i := 0; i < 8; i++ {
		whitePawns[i] = false
		blackPawns[i] = false
		whitePass[i] = false
		blackPass[i] = false
	}
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
	//place the piece to its new target square
	ChessBoard[targetRow][targetCol] = piece

	if whiteTurn == true {
		whiteTurn = false
	} else {
		whiteTurn = true
	}
}

//checks if white pawn move is legal, returns true if legal and false if iillegal
func whitePawnMove(sourceRow int, sourceCol int, targetRow int, targetCol int) bool {
	//moving pawn two squares, pawn should be moving on same file
	if sourceRow-targetRow == 2 && sourceCol == targetCol {

		//if the pawns already moved two squares on their first move then they can't move two squares again
		if whitePawns[sourceCol] == true {
			fmt.Println("You already moved the white pawn two squares.")
			return false
			//checking if any piece blocks the path of the pawn trying to advance two squares
		} else if ChessBoard[sourceRow-1][sourceCol] != "--" || ChessBoard[sourceRow-2][sourceCol] != "--" {

			fmt.Println("There is a piece blocking the white pawn move.")
			return false
			//enabling enpassent for the other player if there is pawn on either side
		}
		if ChessBoard[targetRow][targetCol-1] == "bP" {
			blackPass[targetCol-1] = true
		} else if ChessBoard[targetRow][targetCol+1] == "bP" {
			blackPass[targetCol+1] = true
		}
		//mark the pawn has moved and can't be moved two squares again
		whitePawns[sourceCol] = true

		//moving pawn one square or a pawn capture
	} else if sourceRow-targetRow == 1 {
		
		//determine if its a pawn capture or not, if this is a one square pawn move check if the destination is empty
		if sourceCol == targetCol && ChessBoard[sourceRow-1][sourceCol] == "--" {
			fmt.Println("White Pawn moves one square forward.")
			//mark the pawn has moved and can't be moved two squares
			whitePawns[sourceCol] = true
			//then its a diagonal pawn capture
		} else if (sourceCol-targetCol == 1 || sourceCol-targetCol == -1) && ChessBoard[targetRow][targetCol] != "--" {
			fmt.Println("White pawn captures.")
			//check for enpassent
		} else if whitePass[sourceCol] == true && ChessBoard[targetRow][targetCol] == "--" && (sourceCol-targetCol == 1 || targetCol-sourceCol == 1) {
			//remove black pawn left of white pawn
			ChessBoard[sourceRow][targetCol] = "--"
			fmt.Println("removed black pawn via enpassent")

			//check enpassent the other side now
		} else {
			fmt.Println("Invalid pawn move")
			return false
		}
	} else {
		fmt.Println("Invalid pawn move")
		return false
	}
	//player can only enpassent on the first oppurtunity
	passExpireWhite()
	return true
}

func blackPawnMove(sourceRow int, sourceCol int, targetRow int, targetCol int) bool {
	//moving pawn two squares, pawn should be moving on same file
	if targetRow-sourceRow == 2 && sourceCol == targetCol {

		//if the pawns already moved two squares on their first move then they can't move two squares again
		if blackPawns[sourceCol] == true {
			fmt.Println("You already moved the black pawn two squares.")
			return false
			//checking if any piece blocks the path of the pawn trying to advance two squares
		} else if ChessBoard[sourceRow+1][sourceCol] != "--" || ChessBoard[sourceRow+2][sourceCol] != "--" {

			fmt.Println("There is a piece blocking the black pawn move.")
			return false
		}
		//enabling en passent for other player
		if ChessBoard[targetRow][targetCol-1] == "wP" {
			whitePass[targetCol-1] = true
		} else if ChessBoard[targetRow][targetCol+1] == "wP" {
			whitePass[targetCol+1] = true
		}
		//mark the pawn has moved two squares and can't be moved two squares again
		blackPawns[sourceCol] = true

		//moving pawn one square or a pawn capture
	} else if targetRow-sourceRow == 1 {

		//determine if its a pawn capture or not, if this is a one square pawn move check if the destination is empty
		if sourceCol == targetCol && ChessBoard[sourceRow+1][sourceCol] == "--" {
			fmt.Println("Black Pawn moves one square forward.")
			//mark the pawn has moved and can't be moved two squares
			blackPawns[sourceCol] = true

			//then its a diagonal pawn capture
		} else if (targetCol-sourceCol == 1 || targetCol-sourceCol == -1) && ChessBoard[targetRow][targetCol] != "--" {
			fmt.Println("Black pawn captures.")

		} else if blackPass[sourceCol] == true && ChessBoard[targetRow][targetCol] == "--" && (sourceCol-targetCol == 1 || targetCol-sourceCol == 1) {
			//remove black pawn left of white pawn
			ChessBoard[sourceRow][targetCol] = "--"
			fmt.Println("removed white pawn via enpassent")

		} else {

			fmt.Println("Invalid pawn move")
			return false
		}
	} else {
		fmt.Println("Invalid pawn move")
		return false
	}

	//player can only enpassent on the first oppurtunity
	passExpireBlack()
	return true
}

//enPassent expires for the color if they don't make a move
func passExpireWhite(){
	//setting all the values in the map
	for index, _ := range whitePass{
		whitePass[index] = false
	}
}

func passExpireBlack(){
	//setting all the values to false
	for index, _ := range blackPass{
		blackPass[index] = false
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
