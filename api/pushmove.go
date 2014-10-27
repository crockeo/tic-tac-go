package api

import (
	"github.com/crockeo/tic-tac-go/algorithms"
	"github.com/zenazn/goji/web"
	"net/http"
)

func PushMove(board *algorithms.Board) func(web.C, http.ResponseWriter, *http.Request) {
	return func(c web.C, w http.ResponseWriter, r *http.Request) {

	}
}
