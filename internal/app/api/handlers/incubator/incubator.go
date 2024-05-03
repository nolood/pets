package incubator

import (
	"go.uber.org/zap"
	"pets/internal/app/api/services/incubator"
)

type Handler interface {
}

type incubatorHandler struct {
	s   incubator.Service
	log *zap.Logger
}

func New(log *zap.Logger, s incubator.Service) Handler {
	return &incubatorHandler{s: s, log: log}
}
