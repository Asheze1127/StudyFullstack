package usecase

import "github.com/google/uuid"

// SignInParams サインインパラメータ
type SignInParams struct {
	Email    string
	Password string
}

// SignInOutput サインイン出力
type SignInOutput struct {
	UserID uuid.UUID
	Email  string
	Token  string
}

// SignUpParams サインアップパラメータ
type SignUpParams struct {
	Email    string
	Password string
}

// SignUpOutput サインアップ出力
type SignUpOutput struct {
	UserID uuid.UUID
	Email  string
}
