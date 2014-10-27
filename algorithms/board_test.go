package algorithms

import "testing"

func TestFindWinner(t *testing.T) {
	board1 := Board([][]Player{
		[]Player{X, Nil, Nil},
		[]Player{X, Nil, Nil},
		[]Player{X, Nil, Nil},
	})

	board2 := Board([][]Player{
		[]Player{O, O, O},
		[]Player{Nil, Nil, Nil},
		[]Player{Nil, Nil, Nil},
	})

	board3 := Board([][]Player{
		[]Player{X, Nil, Nil},
		[]Player{Nil, X, Nil},
		[]Player{Nil, Nil, X},
	})

	if FindWinner(board1) != X {
		t.Error("FindWinner on board1 should be X.")
	}

	if FindWinner(board2) != O {
		t.Error("FindWinner on board2 should be O.")
	}

	if FindWinner(board3) != X {
		t.Error("FindWinner on board3 should be X.")
	}
}
