package handlers

import (
	"cyberpets/pets-ws/internal/services/dto"
	"fmt"

	"go.uber.org/zap"
)

type ClickerHandlers interface {
	Handle(data dto.Message)
}

type Handlers struct {
	log *zap.Logger
}

func New(log *zap.Logger) ClickerHandlers {
	return &Handlers{
		log: log,
	}
}

func (h *Handlers) Handle(data dto.Message) {
	const op = "clicker.handlers.handle"

	// newMessage, err := json.Marshal(data)
	// if err != nil {
	// 	h.log.Error(op, zap.Error(err))
	// 	return
	// }

	h.log.Debug(op, zap.String("message", "message received"))

	fmt.Println(data)
}
