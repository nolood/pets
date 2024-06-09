package router

import (
	"errors"
	"fmt"
	"golang.org/x/net/websocket"
	"io"
)

type Service interface {
	Read(conn *websocket.Conn)
}

type routerService struct {
}

func New() Service {
	return &routerService{}
}

func (r *routerService) Read(conn *websocket.Conn) {
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
		_, err = conn.Write([]byte("thank you for the message"))
		if err != nil {
			fmt.Println("write error:", err)
		}
	}
}
