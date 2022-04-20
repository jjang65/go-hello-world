package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"net/http"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,                 // Only server side can access this cookie; no other client side JS can't access
		Path:     "/",                  // entire site
		Secure:   false,                // for now, because we're using http://localhost, must be false; true is for https
		SameSite: http.SameSiteLaxMode, // SameSite Lax mode
	})
	return csrfHandler
}
