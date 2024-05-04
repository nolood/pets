package incubator

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"pets/internal/app/api/services/incubator"
)

type Handler interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type incubatorHandler struct {
	s   incubator.Service
	log *zap.Logger
}

func New(log *zap.Logger, s incubator.Service) Handler {
	return &incubatorHandler{s: s, log: log}
}

func (h *incubatorHandler) Get(w http.ResponseWriter, r *http.Request) {

	h.log.Debug("Incubator get")

	incubatorInst, err := h.s.Get(r.Context())
	if err != nil {
		http.Error(w, "err-incubator-get", http.StatusInternalServerError)
		h.log.Error("Cant get incubator", zap.Error(err))
		return
	}

	response, err := json.Marshal(incubatorInst)
	if err != nil {
		http.Error(w, "err-incubator-marshal", http.StatusInternalServerError)
		h.log.Error("Cant marshal incubator", zap.Error(err))
		return
	}

	_, err = w.Write(response)
	if err != nil {
		http.Error(w, "err-incubator-write", http.StatusInternalServerError)
		h.log.Error("Cant write incubator", zap.Error(err))
	}
}
