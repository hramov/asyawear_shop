package user_repository

import (
	user_dto_out "asyawear/server/internal/app/usecase/user/dto/out"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Repository struct {
	conn *pgxpool.Pool
}

func NewUserRepository(conn *pgxpool.Pool) *Repository {
	return &Repository{
		conn: conn,
	}
}

func (u *Repository) GetUserInfo(ctx context.Context, userId int) (user_dto_out.User, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `SELECT id, email, phone, telegram_username, telegram_id, created_at, updated_at, deleted_at FROM users WHERE id = $1`
	params := []any{userId}
	row := u.conn.QueryRow(ctx, query, params...)

	var res UserModel
	err := row.Scan(&res.Id, &res.Email, &res.Phone, &res.TelegramUsername, &res.TelegramId, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt)
	if err != nil {
		return user_dto_out.User{}, fmt.Errorf("cannot scan user info: %v", err)
	}

	return res.Map(), nil
}
