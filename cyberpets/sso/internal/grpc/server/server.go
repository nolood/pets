package server

import (
	"context"
	ssov1 "cyberpets/protos/gen/go/sso"
	"cyberpets/sso/internal/services/auth"
	"google.golang.org/grpc"
)

type serverAPI struct {
	ssov1.UnimplementedAuthServer
	service auth.Service
}

func Register(gRPC *grpc.Server, service auth.Service) {
	ssov1.RegisterAuthServer(gRPC, &serverAPI{service: service})
}

func (s *serverAPI) Validate(ctx context.Context, req *ssov1.ValidateRequest) (*ssov1.ValidateResponse, error) {
	ok := s.service.Validate(ctx, req)

	return &ssov1.ValidateResponse{Ok: ok}, nil
}

func (s *serverAPI) ValidateToken(ctx context.Context, req *ssov1.ValidateTokenRequest) (*ssov1.ValidateTokenResponse, error) {
	ok, userId := s.service.ValidateToken(ctx, req.Token)

	if !ok {
		return &ssov1.ValidateTokenResponse{Valid: ok, UserId: ""}, nil
	}

	return &ssov1.ValidateTokenResponse{Valid: ok, UserId: userId}, nil
}
