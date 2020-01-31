package server

import (
	"github.com/tmarcus87/rest4grpc/logger"
	"net/http"
	"time"
)

type Server struct {
	native *http.Server
}

func NewServer(addr string, handler http.Handler) *Server {
	return &Server{
		native: &http.Server{
			Addr:         addr,
			Handler:      handler,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
		},
	}
}

func (s *Server) WithTimeout(d time.Duration) *Server {
	s.native.ReadTimeout = d
	s.native.WriteTimeout = d
	return s
}

func (s *Server) Start() error {
	logger.Debugf("Listening server on %s", s.native.Addr)
	return s.native.ListenAndServe()
}

func (s *Server) Close() error {
	return s.native.Close()
}
