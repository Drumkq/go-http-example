package server

import (
	"fmt"
	"log"
	"net/http"

	"example.com/m/internal/config"
	userApi "example.com/m/internal/server/api/v1"
	"github.com/gorilla/mux"
)

type Server struct {
	server *http.Server
	router *mux.Router
}

func New(cfg *config.Config) (Server, error) {
	r := mux.NewRouter()
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:5000",
	}
	outSrv := Server{router: r, server: srv}

	outSrv.initApiV1()

	return outSrv, nil
}

func (s *Server) initApiV1() {
	v1api := s.router.NewRoute().PathPrefix("/v1").Subrouter()

	userApi.New(v1api)
}

func (s *Server) Listen() {
	log.Fatal(s.server.ListenAndServe())
}

func (s *Server) Close() error {
	if err := s.server.Close(); err != nil {
		return fmt.Errorf("error in the server: %w", err)
	}

	return nil
}
