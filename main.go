package main

import (
	"github.com/crockeo/berries"
	"github.com/crockeo/tic-tac-go/api"
	"github.com/crockeo/tic-tac-go/controllers"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web/middleware"
)

// Removing some default middleware.
func removeMiddleware() {
	goji.Abandon(middleware.Recoverer)
	goji.Abandon(middleware.Logger)
}

// Adding my own middleware.
func addMiddleware() {
	goji.Use(berries.StaticMiddleware("static/"))
}

// Registering the routes.
func registerRoutes() {
	// Registering the normal routes.
	goji.Get("/", controllers.GetHome)
	goji.NotFound(controllers.NotFound)

	// Registering the API.
	goji.Get("/api/pull/state", api.PullState)
	goji.Post("/api/push/reset", api.PushReset)
	goji.Post("/api/push/move", api.PushMove)
}

// The entrypoint to the program.
func main() {
	removeMiddleware()
	addMiddleware()
	registerRoutes()

	goji.Serve()
}
