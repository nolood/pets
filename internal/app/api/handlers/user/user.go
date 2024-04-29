package user

import (
	"go.uber.org/zap"
	"net/http"
	"pets/internal/app/api/services/user"
	"pets/internal/domain/models"
)

type Handler interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	service user.Service
	log     *zap.Logger
}

func New(log *zap.Logger, service user.Service) Handler {
	return &userHandler{
		log:     log,
		service: service,
	}
}

func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.User.Create"

	user := models.User{
		Username:  "kek",
		LastName:  "kek",
		FirstName: "kek",
	}

	err := h.service.Create(r.Context(), user)
	if err != nil {
		h.log.Error("HANDLER: create user", zap.Error(err))
	}

	_, err = w.Write([]byte("kek"))
	if err != nil {
		h.log.Error("HANDLER: create user", zap.Error(err))
	}
}
