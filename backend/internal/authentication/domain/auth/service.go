package auth

import (
	"github.com/google/uuid"
)


type Service interface {
	HashPassword(password string) (string, error)

	VerifyPassword(password, hashedPassword string) error

	GenerateToken(userID uuid.UUID, email string) (string, error)

	ValidateToken(token string) (*Claims, error)
}
