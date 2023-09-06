package postgres

import (
	"context"
	"fmt"

	"asyawear/shared/lib/database/types"
	"asyawear/shared/lib/logger"

	"github.com/jackc/pgx/v5/pgxpool"
)

type postgres struct {
	options types.ConnectOptions
	db      *pgxpool.Pool
}

func New(ctx context.Context, cfg types.ConnectOptions, logger logger.Logger) (*pgxpool.Pool, error) {
	pg := &postgres{
		options: cfg,
	}

	err := pg.connect(ctx)
	if err != nil {
		return nil, err
	}

	err = pg.ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot ping postgres connection: %v", err)
	}

	logger.Info(fmt.Sprintf("successfully connected to postgres"))

	return pg.db, nil
}

func (pg *postgres) connect(ctx context.Context) error {
	connPool, err := pgxpool.New(ctx, fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", pg.options.User, pg.options.Password, pg.options.Host, pg.options.Port, pg.options.Database, pg.options.SslMode))
	if err != nil {
		return fmt.Errorf("cannot connect to postgres: %v", err)
	}
	pg.db = connPool
	return nil
}

func (pg *postgres) ping(ctx context.Context) error {
	return pg.db.Ping(ctx)
}

func (pg *postgres) close() {
	pg.db.Close()
}
