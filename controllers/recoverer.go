package controllers

import (
	"github.com/crockeo/berries/tmpl"
	"github.com/zenazn/goji/web"
	"net/http"
)

// Getting the page at home.
func Recoverer(c web.C, w http.ResponseWriter, r *http.Request) {
	err := tmpl.SendPage(c, w, "recoverer", struct{}{})

	if err != nil {
		panic(err)
	}
}
