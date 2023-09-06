package user_dto_out

import "time"

type User struct {
	Id               int       `json:"id"`
	Email            string    `json:"email"`
	Phone            string    `json:"phone"`
	TelegramUsername string    `json:"telegram_username"`
	TelegramId       string    `json:"telegram_id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	DeletedAt        time.Time `json:"deleted_at"`
}
