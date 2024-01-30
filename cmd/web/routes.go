package main

import (
	"github.com/abbassmortazavi/bookings/internal/config"
	"github.com/abbassmortazavi/bookings/internal/handlers"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/generals", handlers.Repo.General)
	mux.Get("/majors", handlers.Repo.Major)

	mux.Get("/search-availability", handlers.Repo.SearchAvailability)
	mux.Post("/search-availability", handlers.Repo.PostSearchAvailability)
	mux.Post("/search-availability-json", handlers.Repo.PostSearchAvailabilityJSON)

	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/make-reservation", handlers.Repo.MakeReservation)
	mux.Post("/make-reservation", handlers.Repo.PostMakeReservation)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}

//5
