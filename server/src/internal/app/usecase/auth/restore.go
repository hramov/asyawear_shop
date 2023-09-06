package auth_usecase

import (
	"asyawear/server/internal/app/usecase"
	"asyawear/server/internal/app/usecase/auth/dto/in"
	auth_dto_out "asyawear/server/internal/app/usecase/auth/dto/out"
	"context"
)

func Restore(ctx context.Context, dto auth_dto_in.Restore, repository AuthRepository, log usecase.Logger) (auth_dto_out.RestoreReply, error) {
	return auth_dto_out.RestoreReply{}, nil
}
