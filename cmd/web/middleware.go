package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
}

// NoSurf adds CSRF protection to all posts requests
func NoSurf(next http.Handler) http.Handler {
	csrHandler := nosurf.New(next)

	csrHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrHandler
}

// SessionLoad loads and saves the session on every requests
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
