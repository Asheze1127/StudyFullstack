package pkg

import "errors"

var (
	// 認証エラー
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrInvalidToken       = errors.New("invalid token")
	ErrTokenExpired       = errors.New("token expired")
	
	// バリデーションエラー
	ErrInvalidInput = errors.New("invalid input")
	
	// 内部エラー
	ErrInternalServer = errors.New("internal server error")
)
