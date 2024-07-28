package main

import (
	"log"
	"net/http"
)

type MiddlewareFunc func(http.Handler) http.Handler

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		log.Println(r.Method)
		next.ServeHTTP(w, r)
	})
}
