package model

import (
	"time"
)

const USER_EMAIL_FIELD = "email"

type User struct {
	BaseModelUuidPk
	Email     string    `gorm:"uniqueIndex;not null;column:email"`      // Email from Cognito
	Name      string    `gorm:"not null;type:varchar(100);column:name"` // User's full name
	LastLogin time.Time `gorm:"nullable;column:last_login"`             // Timestamp for the last login
}
