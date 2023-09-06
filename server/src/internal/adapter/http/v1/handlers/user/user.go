package user_handler

import (
	"asyawear/server/internal/adapter/db/repository/user"
	error_handler "asyawear/server/internal/adapter/http/v1/error"
	"asyawear/server/internal/adapter/http/v1/handlers"
	user_usecase "asyawear/server/internal/app/usecase/user"
	"asyawear/server/internal/config"
	"asyawear/shared/lib/logger"
	"asyawear/shared/lib/utils"
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
	"strconv"
)

type Handler struct {
	handlers.Handler
	Repository user_usecase.UserRepository
}

func New(ctx context.Context, cfg *config.Config, db *pgxpool.Pool, log logger.Logger) *Handler {
	repo := user_repository.NewUserRepository(db)
	return &Handler{Handler: handlers.Handler{Ctx: ctx, Cfg: cfg, Log: log}, Repository: repo}
}

func (h *Handler) Register(r chi.Router) {
	r.Get("/info/{id}", h.info)
}

func (h *Handler) info(w http.ResponseWriter, r *http.Request) {
	userIdRaw := chi.URLParam(r, "id")

	if userIdRaw == "" {
		utils.SendError(utils.StatusBadRequest, "cannot get user id from params", w)
		return
	}

	userId, err := strconv.Atoi(userIdRaw)
	if err != nil {
		utils.SendError(utils.StatusBadRequest, "wrong user id format", w)
		return
	}

	data, err := user_usecase.Info(r.WithContext(h.Ctx).Context(), userId, h.Repository, h.Log)
	if err != nil {
		appErr := error_handler.ParseError(err)
		utils.SendError(appErr.Code, appErr.Message, w)
		return
	}

	utils.SendResponse(utils.StatusOK, data, w)
}
