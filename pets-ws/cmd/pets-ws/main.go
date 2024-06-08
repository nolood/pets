package main

import (
	"cyberpets/logger"
	"cyberpets/pets-ws/internal/config"
	"errors"
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"net/http"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWs(conn *websocket.Conn) {
	fmt.Println("new incoming connection from client:", conn.RemoteAddr())

	s.conns[conn] = true

	s.readLoop(conn)
}

func (s *Server) readLoop(conn *websocket.Conn) {
	buf := make([]byte, 1024)
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
		fmt.Println("read:", string(msg))
		conn.Write([]byte("thank you for the message"))
	}
}

func main() {
	cfg := config.MustLoad()

	log := logger.New(cfg.Env)

	log.Info("starting pets-ws")

	server := NewServer()
	http.Handle("/", websocket.Handler(server.handleWs))
	err := http.ListenAndServe(":5001", nil)
	if err != nil {
		panic(err)
	}
}
