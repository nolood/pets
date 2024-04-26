package logger

import "go.uber.org/zap"

var (
	envLocal = "local"
	envProd  = "production"
)

func New(env string) *zap.Logger {
	var log *zap.Logger

	switch env {
	case envLocal:
		log, _ = zap.NewDevelopment()
	case envProd:
		log, _ = zap.NewProduction()
	default:
		log, _ = zap.NewProduction()
	}

	return log
}
