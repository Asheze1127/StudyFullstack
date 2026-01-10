package auth

import (
	"backend/internal/authentication/domain/auth"
	"github.com/google/uuid"
)

type Client struct {
	passwordService *PasswordService
	jwtService		*JWTService
}

func NewClient() auth.Service {
	return &Client{
		passwordService: NewPasswordService(),
		jwtService:		NewJWTService(),
	}
}

// HashPassword パスワードをハッシュ化
func (c *Client) HashPassword(password string) (string, error) {
	return c.passwordService.HashPassword(password)
}

// VerifyPassword パスワードを検証
func (c *Client) VerifyPassword(hashedPassword, password string) error {
	return c.passwordService.VerifyPassword(hashedPassword, password)
}

// GenerateToken JWTトークンを生成
func (c *Client) GenerateToken(userID uuid.UUID, email string) (string, error) {
	return c.jwtService.GenerateToken(userID, email)
}

// VerifyToken JWTトークンを検証
func (c *Client) ValidateToken(token string) (*auth.Claims, error) {
	return c.jwtService.ValidateToken(token)
}
