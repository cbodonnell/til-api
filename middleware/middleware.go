package middleware

import (
	"net/http"

	auth "github.com/cheebz/go-auth-helpers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func CreateCORSMiddleware(allowedOrigins []string) func(h http.Handler) http.Handler {
	cors := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowCredentials: true,
	})
	return cors.Handler
}

func CreateJwtUUIDMiddleware(authEndpoint string, uuidParam string) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authMap, err := auth.Authenticate(w, r, authEndpoint)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			if s, ok := authMap["uuid"].(string); !ok {
				http.Error(w, "invalid response from auth endpoint", http.StatusUnauthorized)
				return
			} else {
				mux.Vars(r)[uuidParam] = s
				h.ServeHTTP(w, r)
			}
		})
	}
}
