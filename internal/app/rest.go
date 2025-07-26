package app

import (
	"forum/internal/service"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	router  *http.ServeMux
	service service.Service
}

func New(service service.Service) *Handler {
	h := &Handler{
		service: service,
		router:  &http.ServeMux{},
	}

	h.initRoutes()

	return h
}

func (h *Handler) initRoutes() {
	mux := mux.NewRouter()
	h.router.Handle("/api/", mux)
}
