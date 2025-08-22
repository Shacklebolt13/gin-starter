package model

import (
	"time"
)

const USER_EMAIL_FIELD = "email"

type User struct {
	BaseModelStringPk
	Email     string    `gorm:"uniqueIndex;not null;column:email"`
	Name      string    `gorm:"not null;type:varchar(100);column:name"`
	LastLogin time.Time `gorm:"nullable;column:last_login"`
}
