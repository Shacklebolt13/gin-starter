package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Model any

const (
	ID_FIELD         = "id"
	CREATED_AT_FIELD = "created_at"
	UPDATED_AT_FIELD = "updated_at"
	DELETED_AT_FIELD = "deleted_at"
)

type BaseModel struct {
	Model
	CreatedAt time.Time      `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime;column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

type BaseModelIntPk struct {
	ID uint `gorm:"primaryKey;autoIncrement;column:id"`
	BaseModel
}

type BaseModelUuidPk struct {
	ID uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();column:id"`
	BaseModel
}
