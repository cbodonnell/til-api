package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cbodonnell/til-api/middleware"
	"github.com/cbodonnell/til-api/models"
	"github.com/cbodonnell/til-api/services"
	"github.com/gorilla/mux"
)

type MuxHandler struct {
	tilService   services.TilService
	authEndpoint string
	router       *mux.Router
}

type MuxHandlerOptions struct {
	TilService   services.TilService
	AuthEndpoint string
}

// content types
const (
	ContentTypeJSON string = "application/json"
)

var (
	uuidParam = "user_uuid"
)

func NewMuxHandler(opts MuxHandlerOptions) Handler {
	h := &MuxHandler{
		tilService:   opts.TilService,
		authEndpoint: opts.AuthEndpoint,
		router:       mux.NewRouter(),
	}
	h.setupRoutes()
	return h
}

func (h *MuxHandler) setupRoutes() {
	authMiddleware := middleware.CreateJwtUUIDMiddleware(h.authEndpoint, uuidParam)
	auth := h.router.NewRoute().Subrouter() // -> authenticated endpoints
	auth.Use(authMiddleware)
	auth.HandleFunc("/tils", h.GetTilsByUserID).Methods("GET", "OPTIONS")
	auth.HandleFunc("/tils", h.CreateTil).Methods("POST", "OPTIONS")

	sink := h.router.NewRoute().Subrouter() // -> sink to handle all other routes
	sink.PathPrefix("/").HandlerFunc(h.SinkHandler).Methods("GET", "OPTIONS")
}

func (h *MuxHandler) GetRouter() http.Handler {
	return h.router
}

func (h *MuxHandler) AllowOrigins(allowedOrigins []string) {
	cors := middleware.CreateCORSMiddleware(allowedOrigins)
	h.router.Use(cors)
}

func (h *MuxHandler) GetTilsByUserID(w http.ResponseWriter, r *http.Request) {
	user_uuid := mux.Vars(r)[uuidParam]

	users, err := h.tilService.GetAllByUserID(user_uuid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", ContentTypeJSON)
	json.NewEncoder(w).Encode(users)
}

func (h *MuxHandler) CreateTil(w http.ResponseWriter, r *http.Request) {
	user_uuid := mux.Vars(r)[uuidParam]

	var til models.Til
	err := json.NewDecoder(r.Body).Decode(&til)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	til, err = h.tilService.Create(user_uuid, til)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", ContentTypeJSON)
	json.NewEncoder(w).Encode(til)
}

func (h *MuxHandler) SinkHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, fmt.Sprintf("endpoint %s does not exist", r.URL), http.StatusNotFound)
}
