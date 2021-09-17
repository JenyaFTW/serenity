package models

import (
	"github.com/google/uuid"
)

type Project struct {
	ID         uuid.UUID `json:"id" gorm:"default:uuid_generate_v4()"`
	Name       string    `json:"name"`
	MainDomain string    `json:"domain"`
	UserId     uuid.UUID `json:"user_id"`
}
