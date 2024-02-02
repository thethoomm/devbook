package middlewares

import (
	"api/src/authentication"
	"api/src/responses"
	"log"
	"net/http"
)

func Logger(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("| %s %s %s\n", r.Method, r.RequestURI, r.Host)
		nextFunction(w, r)
	}
}

func Authenticate(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := authentication.ValidateToken(r); err != nil {
			responses.Error(w, http.StatusUnauthorized, err)
			return
		}
		nextFunction(w, r)
	}
}
