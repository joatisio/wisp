package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type NullTime sql.NullTime

type ID uuid.UUID

func (id *ID) String() string {
	return uuid.UUID(*id).String()
}

func (id *ID) IsNil() bool {
	return uuid.UUID(*id) == uuid.Nil
}

type BaseModel struct {
	ID        ID        `gorm:"type:uuid;default:uuid_generate_v4()" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt NullTime  `gorm:"index"`
}
