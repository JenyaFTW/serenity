package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Subdomain struct {
	ID         uuid.UUID `json:"id"`
	RootDomain string    `json:"root_domain`
	Value      string    `json:"value"`
	ProjectID  string    `json:"project_id"`
	FirstFound int64     `json:"first_found"`
	LastFound  int64     `json:"last_found"`
}

func (s *Subdomain) BeforeCreate(db *gorm.DB) error {
	s.FirstFound = time.Now().Unix()
	return nil
}
