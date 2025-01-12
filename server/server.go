package server

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

type Server struct {
	Address string
	server  *http.Server
}

func New(address string, handler http.Handler) *Server {
	return &Server{
		Address: address,
		server: &http.Server{
			Addr:         address,
			Handler:      handler,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  30 * time.Second,
		},
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.Address)
	if err != nil {
		return err
	}
	fmt.Println("Listening on", s.Address)

	go handleShutdown(s)

	return s.server.Serve(listener)
}
