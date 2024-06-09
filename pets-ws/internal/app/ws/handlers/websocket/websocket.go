package websockethandler

import (
	"cyberpets/pets-ws/internal/services/router"
	"fmt"
	"golang.org/x/net/websocket"
)

type Server struct {
	connections map[*websocket.Conn]bool
	service     router.Service
}

func New(service router.Service) *Server {
	return &Server{
		service:     service,
		connections: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) Handler(conn *websocket.Conn) {
	fmt.Println(conn.Request().URL.Query())

	s.connections[conn] = true

	s.service.Read(conn)
}
