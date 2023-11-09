package server

import (
	"net/http"

	"github.com/Mario-Benedict/note-api/conf"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	Address string
}

func (s Server) Run() error {
	router := chi.NewRouter()

	conf.SetupRoutes(router)

	return http.ListenAndServe(s.Address, router)
}

func NewServer(address string) *Server {
	return &Server{Address: address}
}
