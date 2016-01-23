package main

import (
	"fmt"
)


//intitalize all pawns to false as they have not moved yet, and also initialize all en passent to false
func initPawns() {
	for i := 0; i < 8; i++ {
		whitePawns[i] = false
		blackPawns[i] = false
		whitePass[i] = false
		blackPass[i] = false
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
		} else if (sourceRow-1 >= 0 && ChessBoard[sourceRow-1][sourceCol] != "--") || (sourceRow-2 >= 0 && ChessBoard[sourceRow-2][sourceCol] != "--") {

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
		if sourceRow-1 >= 0 && sourceCol == targetCol && ChessBoard[sourceRow-1][sourceCol] == "--" {
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
		} else if (sourceRow+1 <= 7 && ChessBoard[sourceRow+1][sourceCol] != "--" ) || (sourceRow+2 <= 7 && ChessBoard[sourceRow+2][sourceCol] != "--") {

			fmt.Println("There is a piece blocking the black pawn move.")
			return false
		}
		//enabling en passent for other player
		if targetCol-1 >= 0 && ChessBoard[targetRow][targetCol-1] == "wP" {
			whitePass[targetCol-1] = true
		} else if targetCol+1<=7 && ChessBoard[targetRow][targetCol+1] == "wP" {
			whitePass[targetCol+1] = true
		}
		//mark the pawn has moved two squares and can't be moved two squares again
		blackPawns[sourceCol] = true

		//moving pawn one square or a pawn capture
	} else if targetRow-sourceRow == 1 {

		//determine if its a pawn capture or not, if this is a one square pawn move check if the destination is empty
		if sourceRow+1 <= 7 && sourceCol == targetCol && ChessBoard[sourceRow+1][sourceCol] == "--" {
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
