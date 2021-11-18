package server

import (
	"backend-task/internal/handlers"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	server *http.Server
}

func New(handler *handlers.Handlers) *Server {
	h := initHandlers(handler)
	g := gin.Default()
	g.POST("/", h.AddTask)
	g.GET("/", h.Tasks)
	return &Server{server: &http.Server{
		Addr:    ":8080",
		Handler: g,
	}}
}

func (s *Server) Run() error {
	return s.server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
