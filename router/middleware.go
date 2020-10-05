package router

import (
	"log"
	"net/http"
)

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL)

		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
