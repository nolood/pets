package router

import (
	"context"
	"cyberpets/pets-ws/internal/app/clicker"
	"cyberpets/pets-ws/internal/domain/models"
	"cyberpets/pets-ws/internal/services/dto"
	"encoding/json"
	"errors"
	"io"

	"go.uber.org/zap"
	"golang.org/x/net/websocket"
)

type Service interface {
	Read(ctx context.Context, conn *websocket.Conn)
}

type routerService struct {
	clicker *clicker.App
	log     *zap.Logger
}

func New(clicker *clicker.App, log *zap.Logger) Service {
	return &routerService{
		clicker: clicker,
		log:     log,
	}
}

func (r *routerService) Read(ctx context.Context, conn *websocket.Conn) {
	const op = "service.router.read"

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			r.log.Error(op, zap.Error(err))
			continue
		}

		var msg dto.Message

		err = json.Unmarshal(buf[:n], &msg)
		if err != nil {
			r.log.Error(op, zap.Error(err))
		}

		switch msg.GameMode {
		case models.ClickerMode:
			r.clicker.Hands.Handle(msg)
		case models.CardMode:
			// r.card.Handle()
		}

		_, err = conn.Write([]byte("thank you for the message " + ctx.Value("user_id").(string)))
		if err != nil {
			r.log.Error(op, zap.Error(err))
		}
	}
}
