package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jojomak13/booking/pkg/handlers"
	"net/http"
)

func routes(handler *handlers.Repository) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handler.Home)
	mux.Get("/about", handler.About)

	return mux
}
