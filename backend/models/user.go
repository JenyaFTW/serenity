package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"default:uuid_generate_v4()"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Avatar   string    `json:"avatarUrl"`
	Role     string    `json:"role" gorm:"default:user"`
}
