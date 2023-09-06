package database_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"asyawear/shared/lib/database/postgres"
	"asyawear/shared/lib/database/types"
	"asyawear/shared/lib/logger"
)

func TestPostgresConnection(t *testing.T) {
	ctx := context.Background()

	l := logger.New("SHARED", logger.Debug, false)

	cfg := types.ConnectOptions{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "postgres",
		Database: "asyawear",
	}

	pg, err := postgres.New(ctx, cfg, l)
	if err != nil {
		t.Error(err.Error())
		return
	}

	row := pg.QueryRow(ctx, "select now()")

	var result time.Time
	row.Scan(&result)
	fmt.Println(result)

	row = pg.QueryRow(ctx, "select now()")

	row.Scan(&result)
	fmt.Println(result)
}
