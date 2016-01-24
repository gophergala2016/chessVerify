package main

import(
	"fmt"
)
func knightMove(sourceRow int, sourceCol int, targetRow int, targetCol int) bool{
	if (targetRow-sourceRow == 2 || targetRow-sourceRow == -2) && (targetCol-sourceCol == 1 || targetCol-sourceCol == -1){
		return true
	}
	return false
}

func bishopMove(sourceRow int, sourceCol int, targetRow int, targetCol int) bool{
	 
	if targetRow < sourceRow && targetCol < sourceCol && sourceRow-targetRow == sourceCol-targetCol{         //left up
    	for i, j := sourceRow-1, sourceCol-1; i > targetRow; i, j = i-1, j-1 {
			if ChessBoard[i][j] != "--"{
				fmt.Println("Piece is blocking the bishop")
				return false
			}
		}    
		            
	}else if targetRow > sourceRow && targetCol < sourceCol && targetRow-sourceRow == sourceCol-targetCol{   //left down	
		for i, j := sourceRow+1, sourceCol-1; i < targetRow; i, j = i+1, j-1 {
			if ChessBoard[i][j] != "--"{
				fmt.Println("Piece is blocking the bishop")
				return false
			}
		}   
	}else if targetRow > sourceRow && targetCol > sourceCol && targetRow-sourceRow == targetCol-sourceCol{    //right down
		for i, j := sourceRow+1, sourceCol+1; i < targetRow; i, j = i+1, j+1 {
			if ChessBoard[i][j] != "--"{
				fmt.Println("Piece is blocking the bishop")
				return false
			}
		}                   
	}else if  targetRow < sourceRow && targetCol > sourceCol && sourceRow-targetRow == targetCol-sourceCol {                                                    //right up
		for i, j := sourceRow-1, sourceCol+1; i < targetRow; i, j = i-1, j+1 {
			if ChessBoard[i][j] != "--"{
				fmt.Println("Piece is blocking the bishop")
				return false
			}
		}                     
	}else{
		fmt.Println("Invalid bishop move")	
		return false
	}
	return true
}

func queenMove(sourceRow int, sourceCol int, targetRow int, targetCol int) bool{
	return true
}

func rookMove(sourceRow int, sourceCol int, targetRow int, targetCol int) bool{
	return true
}

func kingMove(sourceRow int, sourceCol int, targetRow int, targetCol int) bool{
	return true
}