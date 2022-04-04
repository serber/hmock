package server

import (
	"fmt"
)

// Server configuration
type Server struct {
	port uint16
}

func New(port uint16) *Server {
	s := &Server{
		port: port,
	}

	return s
}

func (s Server) PrintInfo() {
	fmt.Println("Port =", s.port)
}
