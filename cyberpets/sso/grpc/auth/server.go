package auth

import (
	"context"
	ssov1 "cyberpets/protos/gen/go/sso"
	"google.golang.org/grpc"
)

type serverAPI struct {
	ssov1.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	ssov1.RegisterAuthServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Validate(ctx context.Context, req *ssov1.ValidateRequest) (*ssov1.ValidateResponse, error) {
	panic("implement me")
	return &ssov1.ValidateResponse{}, nil
}

func (s *serverAPI) ValidateToken(ctx context.Context, req *ssov1.ValidateTokenRequest) (*ssov1.ValidateTokenResponse, error) {
	panic("implement me")
	return &ssov1.ValidateTokenResponse{}, nil
}
