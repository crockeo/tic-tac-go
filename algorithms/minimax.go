package algorithms

import "math"

// Getting the max score for a player.
func maxScore(player Player) int {
	if player == X {
		return -2
	} else if player == O {
		return 2
	} else {
		return 0
	}
}

// Finding the distance between two numbers.
func distance(n1 int, n2 int) int {
	return int(math.Abs(float64(n1 - n2)))
}

// Finding the optimal move for a given player on a given board.
func FindOptimalMove(player Player, board Board) Position {
	return Position{0, 0}
	// TODO: Complete
}
