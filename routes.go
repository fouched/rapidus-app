package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (a *application) routes() *chi.Mux {
	// middleware comes before routes
	a.use(a.Middleware.CheckRemember)

	// routes
	//a.App.Routes.Get("/", a.Handlers.Home)
	//using the convenience wrapper below
	a.get("/", a.Handlers.Home)

	// static routes
	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return a.App.Routes
}
