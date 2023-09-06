package adapter

import (
	v1 "asyawear/server/internal/adapter/http/v1"
	"asyawear/server/internal/config"
	"asyawear/shared/lib/logger"
	"context"
	"fmt"
)

func StartApp(ctx context.Context, cfg *config.Config) {
	log := logger.New(cfg.App.Name, logger.Debug)
	server, err := v1.New(ctx, cfg, log)
	if err != nil {
		log.Error(fmt.Sprintf("cannot start server: %v, exiting...", err))
		return
	}
	server.StartServer(ctx)
}
