package handlers

import (
	"net/http"
)

type Handler interface {
	GetRouter() http.Handler
	AllowOrigins(allowedOrigins []string)
}
