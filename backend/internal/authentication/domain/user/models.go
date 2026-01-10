package user

import (
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
}
