package app

import (
	"log/slog"
	"time"

	grpcApp "github.com/Dasadno/sso/internal/app/grpc"
)

type App struct {
	GRPCrv *grpcApp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	grpcApp := grpcApp.New(log, grpcPort)

	return &App{
		GRPCrv: grpcApp,
	}
}
