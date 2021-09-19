package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Project struct {
	ID         uuid.UUID `json:"id" gorm:"default:uuid_generate_v4()"`
	Name       string    `json:"name"`
	MainDomain string    `json:"domain"`
	UserId     uuid.UUID `json:"user_id"`
	CreatedAt  int64     `json:"created_at"`
}

func (p *Project) BeforeCreate(db *gorm.DB) error {
	p.CreatedAt = time.Now().Unix()
	return nil
}
