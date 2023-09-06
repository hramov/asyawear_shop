package auth

import (
	auth_dto_id "asyawear/server/internal/app/usecase/auth/dto/in"
	user_dto_out "asyawear/server/internal/app/usecase/user/dto/out"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	conn *pgxpool.Pool
}

func NewAuthRepository(conn *pgxpool.Pool) *Repository {
	return &Repository{
		conn: conn,
	}
}

func (u *Repository) GetCandidate(ctx context.Context, dto auth_dto_id.Login) (user_dto_out.User, error) {
	return user_dto_out.User{}, nil
}
