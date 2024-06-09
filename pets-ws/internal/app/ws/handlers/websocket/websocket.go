package websockethandler

import (
	"context"
	jwtclaims "cyberpets/jwt-claims"
	"cyberpets/pets-ws/internal/config"
	"cyberpets/pets-ws/internal/services/router"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
	"golang.org/x/net/websocket"
)

const TOKEN = "token"

type Server struct {
	log         *zap.Logger
	connections map[*websocket.Conn]string
	service     router.Service
	cfg         *config.Config
}

func New(log *zap.Logger, service router.Service, cfg *config.Config) *Server {
	return &Server{
		log:         log,
		service:     service,
		cfg:         cfg,
		connections: make(map[*websocket.Conn]string),
	}
}

// TODO: think about refactor

func (s *Server) Handler(conn *websocket.Conn) {
	query := conn.Request().URL.Query()

	if !query.Has(TOKEN) {
		conn.WriteClose(401)
	}

	tokenString := query.Get(TOKEN)

	var myClaims jwtclaims.Claims
	userID := tokenString

	if s.cfg.Env != "local" {
		token, err := jwt.ParseWithClaims(tokenString, &myClaims, func(token *jwt.Token) (interface{}, error) {
			return []byte(s.cfg.Secret), nil
		})
		if err != nil {
			conn.WriteClose(http.StatusUnauthorized)
		}

		if !token.Valid {
			conn.WriteClose(http.StatusUnauthorized)
		}

		userID = myClaims.Id
	}

	// TODO: fix this
	ctx := context.WithValue(conn.Request().Context(), "user_id", userID)

	s.connections[conn] = userID

	s.service.Read(ctx, conn)
}
