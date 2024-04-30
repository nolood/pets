package user

import (
	"encoding/json"
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
	var entity models.User

	err := json.NewDecoder(r.Body).Decode(&entity)
	if err != nil {
		h.log.Error("Cant decode", zap.Error(err))
		http.Error(w, "err-user-create", http.StatusBadRequest)
		return
	}

	newUser, err := h.service.Create(r.Context(), entity)
	if err != nil {
		h.log.Error("User doesn't created", zap.Error(err))
		http.Error(w, "err-user-create", http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(newUser)
	if err != nil {
		h.log.Error("Cant marshal", zap.Error(err))
		http.Error(w, "err-user-create", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(response)
	if err != nil {
		h.log.Error("Cant write", zap.Error(err))
		http.Error(w, "err-user-create", http.StatusInternalServerError)
	}
}
