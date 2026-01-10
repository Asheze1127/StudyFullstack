package user

import "github.com/google/uuid"

type Repository interface {
	Create(user *User) error
	FindByEmail(email string) (*User, error)
	FindByID(id uuid.UUID) (*User, error)
}
