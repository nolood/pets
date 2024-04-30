package auth

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"pets/internal/app/api/services/auth"
	"pets/internal/domain/models"
)

var (
	ErrLogin = "err-user-login"
	ErrReg   = "err-user-register"
)

type Handler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	service auth.Service
	log     *zap.Logger
}

func New(log *zap.Logger, service auth.Service) Handler {
	return &authHandler{
		log:     log,
		service: service,
	}
}

func (h *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	var entity models.User
	err := json.NewDecoder(r.Body).Decode(&entity)
	if err != nil {
		h.log.Error("Cant decode", zap.Error(err))
		http.Error(w, ErrLogin, http.StatusBadRequest)
		return
	}
}

func (h *authHandler) Register(w http.ResponseWriter, r *http.Request) {
	var entity models.User
	err := json.NewDecoder(r.Body).Decode(&entity)
	if err != nil {
		h.log.Error("Cant decode", zap.Error(err))
		http.Error(w, ErrReg, http.StatusBadRequest)
		return
	}
}
