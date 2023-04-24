package http

import (
	"github.com/gin-gonic/gin"
)

type handler interface {
	Register(r *gin.Engine)
}

type Server struct {
	router  *gin.Engine
	handler handler
}

func NewServer(handler handler) *Server {
	return &Server{router: gin.Default(), handler: handler}
}

func (s *Server) Run(addr string) error {
	s.handler.Register(s.router)

	return s.router.Run(addr)
}
