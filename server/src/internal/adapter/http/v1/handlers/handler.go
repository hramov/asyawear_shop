package handlers

import (
	"asyawear/server/internal/config"
	"asyawear/shared/lib/logger"
	"context"
)

type Handler struct {
	Ctx context.Context
	Cfg *config.Config
	Log logger.Logger
}
