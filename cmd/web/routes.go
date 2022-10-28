package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jojomak13/booking/core/handlers"
	"net/http"
)

func routes(handler *handlers.Repository) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handler.Home)
	mux.Get("/about", handler.About)
	mux.Get("/major", handler.Major)
	mux.Get("/general", handler.General)
	mux.Get("/reservation", handler.Reservation)
	mux.Post("/reservation", handler.Reserve)
	mux.Post("/search", handler.Search)
	mux.Post("/availability", handler.Availability)
	mux.Get("/reservation-summary", handler.ReservationSummary)

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/*", fileServer)

	return mux
}
