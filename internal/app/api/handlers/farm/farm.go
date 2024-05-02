package farm

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"pets/internal/app/api/services/farm"
)

type Handler interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type farmHandler struct {
	service farm.Service
	log     *zap.Logger
}

func New(log *zap.Logger, service farm.Service) Handler {
	return &farmHandler{service: service, log: log}
}

func (h *farmHandler) Get(w http.ResponseWriter, r *http.Request) {
	farmInst, err := h.service.Get(r.Context())
	if err != nil {
		http.Error(w, "err-farm-get", http.StatusInternalServerError)
		h.log.Error("Cant get farm", zap.Error(err))
		return
	}

	response, err := json.Marshal(farmInst)
	if err != nil {
		http.Error(w, "err-farm-get", http.StatusInternalServerError)
		h.log.Error("Cant get farm", zap.Error(err))
		return
	}

	_, err = w.Write(response)
	if err != nil {
		http.Error(w, "err-farm-get", http.StatusInternalServerError)
		h.log.Error("Cant get farm", zap.Error(err))
	}
}
