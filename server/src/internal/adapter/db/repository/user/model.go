package user_repository

import (
	user_dto_out "asyawear/server/internal/app/usecase/user/dto/out"
	"database/sql"
)

type UserModel struct {
	Id               int            `json:"id"`
	Email            sql.NullString `json:"email"`
	Phone            sql.NullString `json:"phone"`
	TelegramUsername sql.NullString `json:"telegram_username"`
	TelegramId       sql.NullString `json:"telegram_id"`
	CreatedAt        sql.NullTime   `json:"created_at"`
	UpdatedAt        sql.NullTime   `json:"updated_at"`
	DeletedAt        sql.NullTime   `json:"deleted_at"`
}

func (u *UserModel) Map() user_dto_out.User {
	return user_dto_out.User{
		Id:               u.Id,
		Email:            u.Email.String,
		Phone:            u.Phone.String,
		TelegramUsername: u.TelegramUsername.String,
		TelegramId:       u.TelegramId.String,
		CreatedAt:        u.CreatedAt.Time,
		UpdatedAt:        u.UpdatedAt.Time,
		DeletedAt:        u.DeletedAt.Time,
	}
}
