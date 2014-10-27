package controllers

import (
	"github.com/crockeo/berries/tmpl"
	"github.com/zenazn/goji/web"
	"log"
	"net/http"
)

// Getting the page at home.
func GetHome(c web.C, w http.ResponseWriter, r *http.Request) {
	err := tmpl.SendPage(c, w, "home", struct{}{})

	if err != nil {
		log.Fatal(err.Error())
	}
}
