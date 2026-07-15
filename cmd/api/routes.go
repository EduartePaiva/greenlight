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
	r.Use(app.authenticate)

	r.NotFound(app.notFoundResponse)
	r.MethodNotAllowed(app.methodNotAllowedResponse)

	r.Get("/v1/healthcheck", app.healthcheckHandler)

	r.With(app.requirePermission("movies:write")).Post("/v1/movies", app.createMovieHandler)
	r.With(app.requirePermission("movies:read")).Get("/v1/movies", app.listMovieHandler)
	r.With(app.requirePermission("movies:read")).Get("/v1/movies/{id}", app.showMovieHandler)
	r.With(app.requirePermission("movies:write")).Patch("/v1/movies/{id}", app.updateMovieHandler)
	r.With(app.requirePermission("movies:write")).Delete("/v1/movies/{id}", app.deleteMovieHandler)

	r.Post("/v1/users", app.registerUserHandler)
	r.Put("/v1/users/activated", app.activateUserHandler)

	r.Post("/v1/tokens/activation", app.createActivationTokenHandler)
	r.Post("/v1/tokens/authentication", app.createAuthenticationTokenHandler)

	return r
}
