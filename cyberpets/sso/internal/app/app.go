package app

import (
	grpcapp "cyberpets/sso/internal/app/grpc"
	"cyberpets/sso/internal/config"

	"go.uber.org/zap"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(log *zap.Logger, cfg *config.Config) *App {

	grpcApp := grpcapp.New(log, cfg)

	return &App{
		GRPCSrv: grpcApp,
	}
}
