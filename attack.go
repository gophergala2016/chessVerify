package main

//returns true if white pawn can attack the specified square
func whitePawnAttack(sourceRow int, sourceCol int, targetRow int, targetCol int) bool {
	if sourceRow-1 == targetRow && sourceCol-1 == targetCol { //capture top left
		return true
	} else if sourceRow-1 == targetRow && sourceCol+1 == targetCol { //capture top right
		return true
	}
	return false
}

func blackPawnAttack(sourceRow int, sourceCol int, targetRow int, targetCol int) bool {
	if sourceRow+1 == targetRow && sourceCol-1 == targetCol { //capture bottom left
		return true
	} else if sourceRow+1 == targetRow && sourceCol+1 == targetCol { //capture bottom right
		return true
	}
	return false
}

func bishopAttack(sourceRow int, sourceCol int, targetRow int, targetCol int) bool {
	for i, j := sourceRow-1, sourceCol-1; i >= targetRow; i, j = i-1, j-1 { //top left diagonal
		if i >= 0 && j >= 0 && ChessBoard[i][j] == "--" {
			if i == targetRow && j == targetCol {
				return true
			}
		} else { //encountered a piece on the diagonal
			if i == targetRow && j == targetCol { //the piece could be our actual target
				return true
			}
			break
		}

	}
	for i, j := sourceRow-1, sourceCol+1; i >= targetRow; i, j = i-1, j+1 { //top right diagonal
		if i >= 0 && j <= 7 && ChessBoard[i][j] == "--" {
			if i == targetRow && j == targetCol {
				return true
			}
		} else { //encountered a piece on the diagonal
			if i == targetRow && j == targetCol { //the piece could be our actual target
				return true
			}
			break
		}

	}
	for i, j := sourceRow+1, sourceCol-1; i <= targetRow; i, j = i+1, j-1 { //bottom left diagonal
		if i <= 7 && j >= 0 && ChessBoard[i][j] == "--" {
			if i == targetRow && j == targetCol {
				return true
			}
		} else { //encountered a piece on the diagonal
			if i == targetRow && j == targetCol { //the piece could be our actual target
				return true
			}
			break
		}

	}
	for i, j := sourceRow+1, sourceCol+1; i <= targetRow; i, j = i+1, j+1 { //bottom right diagonal
		if i <= 7 && j <= 7 && ChessBoard[i][j] == "--" {
			if i == targetRow && j == targetCol {
				return true
			}
		} else { //encountered a piece on the diagonal
			if i == targetRow && j == targetCol { //the piece could be our actual target
				return true
			}
			break
		}

	}

	return false
}

func knightAttack(sourceRow int, sourceCol int, targetRow int, targetCol int) bool {
	if (targetRow-sourceRow == 2 && targetCol-sourceCol == 1) || (targetRow-sourceRow == -2 && targetCol-sourceCol == -1) {
		return true
	} else if (targetRow-sourceRow == 2 && targetCol-sourceCol == -1) || (targetRow-sourceRow == -2 && targetCol-sourceCol == 1) {
		return true
	} else if (targetRow-sourceRow == 1 && targetRow-sourceRow == 2) || (targetCol-sourceCol == 1 && targetCol-sourceCol == -2) {
		return true
	} else if (targetRow-sourceRow == -1 && targetRow-sourceRow == 2) || (targetCol-sourceCol == -1 && targetCol-sourceCol == -2) {
		return true
	}
	return false
}

//rook + bishop movments
func queenAttack(sourceRow int, sourceCol int, targetRow int, targetCol int) bool {
	//bishop movements
	for i, j := sourceRow-1, sourceCol-1; i >= targetRow; i, j = i-1, j-1 { //top left diagonal
		if i >= 0 && j >= 0 && ChessBoard[i][j] == "--" {
			if i == targetRow && j == targetCol {
				return true
			}
		} else { //encountered a piece on the diagonal
			if i == targetRow && j == targetCol { //the piece could be our actual target
				return true
			}
			break
		}

	}
	for i, j := sourceRow-1, sourceCol+1; i >= targetRow; i, j = i-1, j+1 { //top right diagonal
		if i >= 0 && j <= 7 && ChessBoard[i][j] == "--" {
			if i == targetRow && j == targetCol {
				return true
			}
		} else { //encountered a piece on the diagonal
			if i == targetRow && j == targetCol { //the piece could be our actual target
				return true
			}
			break
		}

	}
	for i, j := sourceRow+1, sourceCol-1; i <= targetRow; i, j = i+1, j-1 { //bottom left diagonal
		if i <= 7 && j >= 0 && ChessBoard[i][j] == "--" {
			if i == targetRow && j == targetCol {
				return true
			}
		} else { //encountered a piece on the diagonal
			if i == targetRow && j == targetCol { //the piece could be our actual target
				return true
			}
			break
		}

	}
	for i, j := sourceRow+1, sourceCol+1; i <= targetRow; i, j = i+1, j+1 { //bottom right diagonal
		if i <= 7 && j <= 7 && ChessBoard[i][j] == "--" {
			if i == targetRow && j == targetCol {
				return true
			}
		} else { //encountered a piece on the diagonal
			if i == targetRow && j == targetCol { //the piece could be our actual target
				return true
			}
			break
		}

	}
	//rook movements
	for i := sourceRow + 1; i <= targetRow; i++ { //up
		if i <= 7 && ChessBoard[i][sourceCol] == "--" {
			if i == targetRow && sourceCol == targetCol {
				return true
			}
		} else {
			if i == targetRow && sourceCol == targetCol {
				return true
			}
			break
		}

	}
	for i := sourceRow - 1; i >= targetRow; i-- { //down
		if i >= 0 && ChessBoard[i][sourceCol] == "--" {
			if i == targetRow && sourceCol == targetCol {
				return true
			}
		} else {
			if i == targetRow && sourceCol == targetCol {
				return true
			}
			break
		}

	}
	for i := sourceCol + 1; i <= targetCol; i++ { //right
		if i <= 7 && ChessBoard[sourceRow][i] == "--" {
			if i == targetCol && sourceRow == targetRow {
				return true
			}
		} else {
			if i == targetCol && sourceRow == targetRow {
				return true
			}
			break
		}

	}
	for i := sourceCol - 1; i >= targetCol; i-- { //left
		if i >= 0 && ChessBoard[sourceRow][i] == "--" {
			if i == targetCol && sourceRow == targetRow {
				return true
			}
		} else {
			if i == targetCol && sourceRow == targetRow {
				return true
			}
			break
		}

	}

	return false
}

func rookAttack(sourceRow int, sourceCol int, targetRow int, targetCol int) bool {
	for i := sourceRow + 1; i <= targetRow; i++ { //up
		if i <= 7 && ChessBoard[i][sourceCol] == "--" {
			if i == targetRow && sourceCol == targetCol {
				return true
			}
		} else {
			if i == targetRow && sourceCol == targetCol {
				return true
			}
			break
		}

	}
	for i := sourceRow - 1; i >= targetRow; i-- { //down
		if i >= 0 && ChessBoard[i][sourceCol] == "--" {
			if i == targetRow && sourceCol == targetCol {
				return true
			}
		} else {
			if i == targetRow && sourceCol == targetCol {
				return true
			}
			break
		}

	}
	for i := sourceCol + 1; i <= targetCol; i++ { //right
		if i <= 7 && ChessBoard[sourceRow][i] == "--" {
			if i == targetCol && sourceRow == targetRow {
				return true
			}
		} else {
			if i == targetCol && sourceRow == targetRow {
				return true
			}
			break
		}

	}
	for i := sourceCol - 1; i >= targetCol; i-- { //left
		if i >= 0 && ChessBoard[sourceRow][i] == "--" {
			if i == targetCol && sourceRow == targetRow {
				return true
			}
		} else {
			if i == targetCol && sourceRow == targetRow {
				return true
			}
			break
		}

	}
	return false
}

func kingAttack(sourceRow int, sourceCol int, targetRow int, targetCol int) bool {

	if (sourceRow-targetRow == 1 || sourceRow-targetRow == -1) && (sourceCol-targetCol == -1 || sourceCol-targetCol == 1) { // four diagonals
		return true
	} else if (sourceRow-targetRow == 1 || sourceRow-targetRow == -1) && sourceCol == targetCol { //up or down
		return true
	} else if (sourceCol-targetCol == 1 || sourceCol-targetCol == -1) && sourceRow == targetRow { //left or right
		return true
	}
	return false
}
