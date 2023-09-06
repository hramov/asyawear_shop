package user_usecase

import (
	"asyawear/server/internal/app/usecase"
	user_dto_out "asyawear/server/internal/app/usecase/user/dto/out"
	"context"
)

func Info(ctx context.Context, userId int, repository UserRepository, log usecase.Logger) (user_dto_out.User, error) {
	return repository.GetUserInfo(ctx, userId)
}
