package algorithms

// A type to represent a position on the board.
type Position struct {
	Row int
	Col int
}

// A type to distinguish between a normal int and one that represents a player.
type Player int

// An enum to represent the different states on the board.
const (
	X   Player = 1  // Representing the X-player.
	O   Player = -1 // Representing the Y-player
	Nil Player = 0  // Representing no player.
)

// Calculating the score that means a player has won.
func WinScore(player Player) int {
	return int(player * 3)
}

// A type to represent a board.
type Board [][]Player

// Getting the other player.
func Other(player Player) Player {
	if player == X {
		return O
	} else if player == O {
		return X
	} else {
		return Nil
	}
}

// Getting the value at a position in a board.
func BoardPull(pos Position, board Board) Player {
	return board[pos.Row][pos.Col]
}

// Setting the value at a position in a board.
func BoardPush(pos Position, state Player, board Board) Board {
	board[pos.Row][pos.Col] = state
	return board
}

// Checking if a move is valid.
func IsMoveValid(pos Position, board Board) bool {
	return BoardPull(pos, board) == Nil
}

// Finding each valid move in a board.
func FindValidMoves(board Board) []Position {
	poss := []Position{}
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			pos := Position{row, col}
			if IsMoveValid(pos, board) {
				poss = append(poss, pos)
			}
		}
	}

	return poss
}

// Finding a winner -- if there is one.
func FindWinner(board Board) Player {
	for row := 0; row < 3; row++ {
		sum := 0
		for col := 0; col < 3; col++ {
			sum += int(BoardPull(Position{row, col}, board))
		}

		if sum == WinScore(X) {
			return X
		} else if sum == WinScore(O) {
			return O
		}
	}

	for col := 0; col < 3; col++ {
		sum := 0
		for row := 0; row < 3; row++ {
			sum += int(BoardPull(Position{row, col}, board))
		}

		if sum == WinScore(X) {
			return X
		} else if sum == WinScore(O) {
			return O
		}
	}

	diags := [][]Position{
		[]Position{
			Position{0, 0},
			Position{1, 1},
			Position{2, 2},
		},

		[]Position{
			Position{0, 2},
			Position{1, 1},
			Position{2, 0},
		},
	}

	for i := 0; i < len(diags); i++ {
		sum := 0
		for p := 0; p < len(diags[i]); p++ {
			sum += int(BoardPull(diags[i][p], board))
		}

		if sum == WinScore(X) {
			return X
		} else if sum == WinScore(O) {
			return O
		}
	}

	return Nil
}

// Checking if the game is over.
func IsOver(board Board) bool {
	return len(FindValidMoves(board)) == 0 || FindWinner(board) != Nil
}
