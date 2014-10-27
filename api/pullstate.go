package api

import (
	"encoding/json"
	"github.com/crockeo/tic-tac-go/algorithms"
	"github.com/zenazn/goji/web"
	"net/http"
)

// A type to more easily be used in the PullState function.
type pullStateResponse struct {
	Over   bool  `json:"over"`
	Winner int   `json:"winner"`
	Turn   int   `json:"turn"`
	State  []int `json:"state"`
}

// Flattening a board.
func flatten(board algorithms.Board) []int {
	ints := make([]int, 9)

	i := 0
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			ints[i] = int(algorithms.BoardPull(algorithms.Position{row, col}, board))
			i++
		}
	}

	return ints
}

func PullState(board *algorithms.Board) func(web.C, http.ResponseWriter, *http.Request) {
	return func(c web.C, w http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(w)

		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		err := enc.Encode(pullStateResponse{
			algorithms.IsOver(*board),
			int(algorithms.FindWinner(*board)),
			int(algorithms.FindTurn(*board)),
			flatten(*board),
		})

		if err != nil {
			panic(err)
		}
	}
}
