package usecase

import (
	"backend/internal/authentication/domain/auth"
	"backend/internal/authentication/domain/user"
	"backend/pkg"
	"errors"
)

// SignInUseCase サインインユースケースインターフェース
type SignInUseCase interface {
	Execute(params *SignInParams) (*SignInOutput, error)
}

type signInUseCase struct {
	userRepo user.Repository
	authSvc  auth.Service
}

// NewSignInUseCase コンストラクタ
func NewSignInUseCase(
	userRepo user.Repository,
	authSvc auth.Service,
) SignInUseCase {
	return &signInUseCase{
		userRepo: userRepo,
		authSvc:  authSvc,
	}
}

// Execute サインインを実行
func (uc *signInUseCase) Execute(params *SignInParams) (*SignInOutput, error) {
	// ユーザーを検索
	u, err := uc.userRepo.FindByEmail(params.Email)
	if err != nil {
		return nil, pkg.ErrInvalidCredentials
	}
	
	// パスワード検証
	if err := uc.authSvc.VerifyPassword(u.Password, params.Password); err != nil {
		return nil, pkg.ErrInvalidCredentials
	}
	
	// JWTトークン生成
	token, err := uc.authSvc.GenerateToken(u.ID, u.Email)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}
	
	return &SignInOutput{
		UserID: u.ID,
		Email:  u.Email,
		Token:  token,
	}, nil
}