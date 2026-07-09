package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()

	r.Use(app.recoverPanic)
	if app.config.limiter.enabled {
		r.Use(app.rateLimit)
	}

	r.NotFound(app.notFoundResponse)
	r.MethodNotAllowed(app.methodNotAllowedResponse)

	r.Get("/v1/healthcheck", app.healthcheckHandler)

	r.Post("/v1/movies", app.createMovieHandler)
	r.Get("/v1/movies", app.listMovieHandler)
	r.Get("/v1/movies/{id}", app.showMovieHandler)
	r.Patch("/v1/movies/{id}", app.updateMovieHandler)
	r.Delete("/v1/movies/{id}", app.deleteMovieHandler)

	r.Post("/v1/users", app.registerUserHandler)
	r.Put("/v1/users/activated", app.activateUserHandler)

	r.Post("/v1/tokens/activation", app.createActivationTokenHandler)

	return r
}
