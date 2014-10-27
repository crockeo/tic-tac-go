package main

import (
	"github.com/crockeo/berries"
	"github.com/crockeo/tic-tac-go/algorithms"
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
	goji.Use(berries.RecoveryMiddleware(controllers.Recoverer))
}

// Registering the routes.
func registerRoutes() {
	// Registering the normal routes.
	goji.Get("/", controllers.GetHome)
	goji.NotFound(controllers.NotFound)

	board := &algorithms.EmptyBoard

	// Registering the API.
	goji.Get("/api/pull/state", api.PullState(board))
	goji.Post("/api/push/reset", api.PushReset(board))
	goji.Post("/api/push/move", api.PushMove(board))
}

// The entrypoint to the program.
func main() {
	removeMiddleware()
	addMiddleware()
	registerRoutes()

	goji.Serve()
}
