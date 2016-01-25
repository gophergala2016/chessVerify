package main

//can't castle in check or through check, can't castle when king has already moved
func canWhiteCastleKing() bool{
	
	if wkMoved == false && wkrMoved == false && canBlackKillSquare(7, 4) == false && canBlackKillSquare(7, 5) == false && canBlackKillSquare(7, 6) == false{
		return true
	}
	return false
}

func canBlackCastleKing() bool{
	if bkMoved == false && bkrMoved == false && canWhiteKillSquare(0, 4) == false && canBlackKillSquare(0, 5) == false && canBlackKillSquare(0, 6) == false{
		return true
	}
	return false
}

func canWhiteCastleQueen() bool{
	if wkMoved == false && wqrMoved == false && canBlackKillSquare(7, 3) == false && canBlackKillSquare(7, 2) == false && canBlackKillSquare(7, 1) == false{
		return true
	}
	return false
}

func canBlackCastleQueen() bool{
	if bkMoved == false && bqrMoved == false && canWhiteKillSquare(0, 3) == false && canWhiteKillSquare(0, 2) == false && canWhiteKillSquare(0, 1) == false{
		return true
	}
	return false
}
