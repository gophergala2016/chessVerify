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
	[]string{"-", "-", "-", "-", "-", "-", "-", "-"},
	[]string{"-", "-", "-", "-", "-", "-", "-", "-"},
	[]string{"-", "-", "-", "-", "-", "-", "-", "-"},
	[]string{"-", "-", "-", "-", "-", "-", "-", "-"},
	[]string{"wP", "wP", "wP", "wP", "wP", "wP", "wP", "wP"},
	[]string{"wR", "wN", "wB", "wQ", "wK", "wB", "wN", "wR"},
}

//stores array of booleans indicating whether or not the pawns have made their first move yet
var whitePawns [8]bool
var blackPawns[8]bool

func main() {
	initPawns()
	printBoard()
	result := chessVerify("e2", "e4")
	
	if result == true{
		fmt.Println("Valid move")
		printBoard()
		
	}else{
		fmt.Println("Invalid move")
	}
}

//returns true if the move is valid otherwise it returns false
func chessVerify(source string, target string) bool{
	var sourceCol = source[0]
	//if a white piece is picked and its blacks turn or if a black piece is picked and its whites turn or if no piece is picked up return false
	if (whiteTurn == true && sourceCol == 'b') || (whiteTurn == false && sourceCol == 'w') || source == " --"{
		return false
	}
		
	
	var sourceRow = source[1]
	fmt.Printf("%c %c\n", sourceCol, sourceRow)
	var targetCol = target[0]
	var targetRow = target[1]
	
	//converting to proper format
	newSourceCol := 8-convertLetter(sourceCol)
	newSourceRow := 8-int(sourceRow-'0')
	
	newTargetCol := 8-convertLetter(targetCol)
	newTargetRow := 8-int(targetRow-'0')
	
	fmt.Printf("%s\n", ChessBoard[newSourceRow][newSourceCol])
	fmt.Printf("%s\n", ChessBoard[newTargetRow][newTargetCol])
	
	makeMove(newSourceRow, newSourceCol, newTargetRow, newTargetCol, source)
	
	return true
}

//intitalize all pawns to false as they have not moved yet
func initPawns(){
	for i:=0; i<8; i++{
		whitePawns[i] = false
		blackPawns[i] = false
	}
}

//changes chess letter notation to a number a=0, b=1, c=2, etc
func convertLetter(letter byte) int{
	switch letter{
		case 'a':
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
		default:         //this is h on the chess board
			return 7
	}
	
			
}

//makes the chess move on the board, verify the move first
func makeMove(sourceRow int, sourceCol int, targetRow int, targetCol int, source string){
	
	//make the source square blank as now the piece is no longer there
	ChessBoard[sourceRow][sourceCol] = "-"
	//place the piece to its new target square
	ChessBoard[targetRow][targetCol] = source
	
	
	if whiteTurn == true{
		whiteTurn = false
	}else{
		whiteTurn = true
	}
}

//checks if white pawn move is legal
func whitepPawnMove(){
	
}

func blackPawnMove(){
	
}

func printBoard() {

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Printf("%s ", ChessBoard[i][j])
		}
		fmt.Printf("\n")
	}
}