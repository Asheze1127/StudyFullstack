package usecase

import (
	"backend/internal/authentication/domain/auth"
	"backend/internal/authentication/domain/user"
	"backend/pkg"
	"errors"
)

// SignUpUseCase サインアップユースケースインターフェース
type SignUpUseCase interface {
	Execute(params *SignUpParams) (*SignUpOutput, error)
}

type signUpUseCase struct {
	userRepo user.Repository
	authSvc  auth.Service
}

// NewSignUpUseCase コンストラクタ
func NewSignUpUseCase(
	userRepo user.Repository,
	authSvc auth.Service,
) SignUpUseCase {
	return &signUpUseCase{
		userRepo: userRepo,
		authSvc:  authSvc,
	}
}

// Execute サインアップを実行
func (uc *signUpUseCase) Execute(params *SignUpParams) (*SignUpOutput, error) {
	// 既存ユーザーチェック
	_, err := uc.userRepo.FindByEmail(params.Email)
	if err == nil {
		// ユーザーが既に存在する
		return nil, pkg.ErrEmailAlreadyExists
	}
	
	// パスワードハッシュ化
	hashedPassword, err := uc.authSvc.HashPassword(params.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}
	
	// ユーザー作成
	newUser := &user.User{
		Email:    params.Email,
		Password: hashedPassword,
	}
	
	if err := uc.userRepo.Create(newUser); err != nil {
		return nil, errors.New("failed to create user")
	}
	
	return &SignUpOutput{
		UserID: newUser.ID,
		Email:  newUser.Email,
	}, nil
}
