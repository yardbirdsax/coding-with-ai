package server

import (
	"context"
	"log"
	"net/http"
	"time"
)

// Server represents the HTTP server
type Server struct {
	server *http.Server
	logger *log.Logger
}

// New creates a new server instance
func New(addr string, logger *log.Logger) *Server {
	return &Server{
		server: &http.Server{
			Addr:    addr,
			Handler: http.NewServeMux(),
		},
		logger: logger,
	}
}

// RegisterHandler registers a handler for a specific path
func (s *Server) RegisterHandler(path string, handler http.Handler) {
	mux := s.server.Handler.(*http.ServeMux)
	mux.Handle(path, handler)
}

// Start starts the HTTP server
func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

// Stop gracefully shuts down the server
func (s *Server) Stop(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	
	return s.server.Shutdown(ctx)
}