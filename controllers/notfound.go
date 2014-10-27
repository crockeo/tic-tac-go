package controllers

import (
	"github.com/crockeo/berries/tmpl"
	"github.com/zenazn/goji/web"
	"net/http"
)

func NotFound(c web.C, w http.ResponseWriter, r *http.Request) {
	err := tmpl.SendPage(c, w, "notfound", struct{}{})

	if err != nil {
		panic(err)
	}
}
