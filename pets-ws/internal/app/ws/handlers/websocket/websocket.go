package websockethandler

import (
	"errors"
	"fmt"
	"golang.org/x/net/websocket"
	"io"
)

type Server struct {
	connections map[*websocket.Conn]bool
}

func New() *Server {
	return &Server{
		connections: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) Handler(conn *websocket.Conn) {
	fmt.Println("new incoming connection from client:", conn.RemoteAddr())

	fmt.Println(conn.Request().URL.Query())

	// TODO: think about s.connections[conn] = userId + need to validate
	// TODO: validate move to service layer
	s.connections[conn] = true

	s.readLoop(conn)
}

// readLoop TODO: move to service layer
func (s *Server) readLoop(conn *websocket.Conn) {
	buf := make([]byte, 1024)
	clientAddr := s.connections[conn]
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			fmt.Println("read error:", err)
			continue
		}

		msg := buf[:n]
		fmt.Println("read:", string(msg), "from", clientAddr)
		conn.Write([]byte("thank you for the message"))
	}
}
