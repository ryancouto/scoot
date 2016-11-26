package server

import (
	"fmt"
	"log"
)

type Addr string

type Server struct {
	addr Addr
}

func NewServer(addr Addr) *Server {
	return &Server{addr}
}

func (s *Server) Serve() error {
	log.Println("Hello world", s.addr)
	return fmt.Errof("not yet implemented")
}
