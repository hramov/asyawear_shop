package v1

import (
	auth_handler "asyawear/server/internal/adapter/http/v1/handlers/auth"
	user_handler "asyawear/server/internal/adapter/http/v1/handlers/user"
	"asyawear/server/internal/config"
	"asyawear/shared/lib/database/postgres"
	"asyawear/shared/lib/database/types"
	"asyawear/shared/lib/logger"
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
	"time"
)

type Api struct {
	ctx    context.Context
	config *config.Config
	log    logger.Logger
	db     *pgxpool.Pool
}

func New(ctx context.Context, cfg *config.Config, log logger.Logger) (*Api, error) {
	connectOptions := types.ConnectOptions{
		Host:            cfg.Db.Postgres.Host,
		Port:            cfg.Db.Postgres.Port,
		User:            cfg.Db.Postgres.User,
		Password:        cfg.Db.Postgres.Password,
		Database:        cfg.Db.Postgres.Database,
		SslMode:         "disable",
		MaxOpenCons:     50,
		MaxIdleCons:     10,
		ConnMaxIdleTime: 1 * time.Minute,
		ConnMaxLifetime: 5 * time.Minute,
	}

	db, err := postgres.New(ctx, connectOptions, log)
	if err != nil {
		return nil, err
	}
	return &Api{
		ctx:    ctx,
		config: cfg,
		log:    log,
		db:     db,
	}, nil
}

func (a *Api) registerHandlers(r chi.Router) {

	auth := auth_handler.New(a.ctx, a.config, a.db, a.log)
	r.Route("/auth", auth.Register)

	user := user_handler.New(a.ctx, a.config, a.db, a.log)
	r.Route("/user", user.Register)
}

func (a *Api) StartServer(ctx context.Context) {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	r.Route("/api", a.registerHandlers)

	go func() {
		a.log.Info("starting server")
		if err := http.ListenAndServe(fmt.Sprintf(":%d", a.config.App.Port), r); err != nil {
			a.log.Error(fmt.Sprintf("cannot start server: %v", err))
			return
		}
	}()

	<-ctx.Done()
	a.log.Info(fmt.Sprintf("starting graceful shutdown for server"))
	err := a.StopServer()
	if err != nil {
		a.log.Error(fmt.Sprintf("cannot stop server"))
	}
}

func (a *Api) StopServer() error {
	a.db.Close()
	return nil
}
