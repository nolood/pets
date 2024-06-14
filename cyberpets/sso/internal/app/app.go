package app

import (
	grpcapp "cyberpets/sso/internal/app/grpc"
	"cyberpets/sso/internal/config"

	"go.uber.org/zap"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(log *zap.Logger, grpcCfg config.GRPC) *App {

	grpcApp := grpcapp.New(log, grpcCfg.Port)

	return &App{
		GRPCSrv: grpcApp,
	}
}
