package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	ctx        context.Context
	port       int
	router     http.Handler
	httpServer *http.Server
}

func NewServer(ctx context.Context, port int) (*Server, error) {

	handler := NewHandler(ctx)

	router := registerRoutes(handler)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	return &Server{
		ctx:        ctx,
		router:     router,
		port:       port,
		httpServer: server,
	}, nil
}

func (s *Server) Start() error {
	log.Printf("Rest API gateway started on port: %d", s.port)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
