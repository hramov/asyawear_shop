package auth_handler

import (
	"asyawear/server/internal/adapter/db/repository/auth"
	error_handler "asyawear/server/internal/adapter/http/v1/error"
	"asyawear/server/internal/adapter/http/v1/handlers"
	auth_usecase "asyawear/server/internal/app/usecase/auth"
	auth_dto_id "asyawear/server/internal/app/usecase/auth/dto/in"
	"asyawear/server/internal/config"
	"asyawear/shared/lib/logger"
	"asyawear/shared/lib/utils"
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

type Handler struct {
	handlers.Handler
	Repository auth_usecase.AuthRepository
}

func New(ctx context.Context, cfg *config.Config, db *pgxpool.Pool, log logger.Logger) *Handler {
	repo := auth.NewAuthRepository(db)
	return &Handler{Handler: handlers.Handler{Ctx: ctx, Cfg: cfg, Log: log}, Repository: repo}
}

func (h *Handler) Register(r chi.Router) {
	r.Post("/login", h.login)
	r.Post("/register", h.register)
	r.Post("/restore", h.restore)
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	dto, err := utils.GetBody[auth_dto_id.Login](r)
	if err != nil {
		utils.SendError(utils.StatusInternalServerError, "cannot get request body", w)
		return
	}

	data, err := auth_usecase.Login(r.WithContext(h.Ctx).Context(), dto, h.Repository, h.Log)
	if err != nil {
		appErr := error_handler.ParseError(err)
		utils.SendError(appErr.Code, appErr.Message, w)
		return
	}

	utils.SendResponse(utils.StatusOK, data, w)
}

func (h *Handler) register(w http.ResponseWriter, r *http.Request) {
	dto, err := utils.GetBody[auth_dto_id.Register](r)
	if err != nil {
		utils.SendError(utils.StatusInternalServerError, "cannot get request body", w)
		return
	}

	data, err := auth_usecase.Register(r.WithContext(h.Ctx).Context(), dto, h.Repository, h.Log)
	if err != nil {
		appErr := error_handler.ParseError(err)
		utils.SendError(appErr.Code, appErr.Message, w)
		return
	}

	utils.SendResponse(utils.StatusOK, data, w)
}

func (h *Handler) restore(w http.ResponseWriter, r *http.Request) {
	dto, err := utils.GetBody[auth_dto_id.Restore](r)
	if err != nil {
		utils.SendError(utils.StatusInternalServerError, "cannot get request body", w)
		return
	}

	data, err := auth_usecase.Restore(r.WithContext(h.Ctx).Context(), dto, h.Repository, h.Log)
	if err != nil {
		appErr := error_handler.ParseError(err)
		utils.SendError(appErr.Code, appErr.Message, w)
		return
	}

	utils.SendResponse(utils.StatusOK, data, w)
}
