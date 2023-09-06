package user_usecase

import (
	user_dto_out "asyawear/server/internal/app/usecase/user/dto/out"
	"context"
)

type UserRepository interface {
	GetUserInfo(ctx context.Context, userId int) (user_dto_out.User, error)
}
