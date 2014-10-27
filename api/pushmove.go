package api

import (
	"github.com/crockeo/berries/json"
	"github.com/crockeo/tic-tac-go/algorithms"
	"github.com/zenazn/goji/web"
	"io/ioutil"
	"net/http"
)

// A type that represents the request from the user.
type pushMoveRequest struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

// The response to a PushMove.
type pushMoveResponse struct {
	Error   bool   `json:"error"`
	Msg     string `json:"msg"`
	Refresh bool   `json:"refresh"`
}

func PushMove(board *algorithms.Board) func(web.C, http.ResponseWriter, *http.Request) {
	return func(c web.C, w http.ResponseWriter, r *http.Request) {
		bs, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		var pmr pushMoveRequest
		json.UDecode(&pmr, string(bs))

		*board = algorithms.BoardPush(algorithms.Position{pmr.Row, pmr.Col}, algorithms.X, *board)
	}
}
