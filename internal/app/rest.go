package app

import (
	"forum/internal/service"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	router  *http.ServeMux
	service *service.Service
}

func New(service *service.Service) *Handler {
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

	mux.HandleFunc("/api/forum/create", h.CreateForum).Methods("POST")
	mux.HandleFunc("/api/user/{nickname}/create", h.CreateUser).Methods("POST")
	mux.HandleFunc("/api/user/{nickname}/profile", h.GetUser).Methods("GET")
	mux.HandleFunc("/api/user/{nickname}/profile", h.ChangeUserProfile).Methods("POST")
}

func (h *Handler) Router() *http.ServeMux {
	return h.router
}
