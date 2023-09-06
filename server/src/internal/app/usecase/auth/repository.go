package auth_usecase

import (
	auth_dto_id "asyawear/server/internal/app/usecase/auth/dto/in"
	user_dto_out "asyawear/server/internal/app/usecase/user/dto/out"
	"context"
)

type AuthRepository interface {
	GetCandidate(ctx context.Context, dto auth_dto_id.Login) (user_dto_out.User, error)
}
