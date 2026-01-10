package user

import (
	"backend/internal/authentication/domain/user"
	"backend/pkg"
	"errors"
	
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Client ユーザーサービス実装（GORMを使用、Prismaはマイグレーション管理のみ）
type Client struct {
	db *gorm.DB
}

// NewClient コンストラクタ（GORMのDBインスタンスを受け取る）
func NewClient(db *gorm.DB) user.Repository {
	return &Client{
		db: db,
	}
}

// Create ユーザーを作成
func (c *Client) Create(u *user.User) error {
	if err := c.db.Create(u).Error; err != nil {
		// 重複エラーのチェック（PostgreSQLの場合）
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return pkg.ErrEmailAlreadyExists
		}
		return err
	}
	return nil
}

// FindByEmail メールアドレスでユーザーを検索
func (c *Client) FindByEmail(email string) (*user.User, error) {
	var u user.User
	if err := c.db.Where("email = ?", email).First(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, pkg.ErrInvalidCredentials
		}
		return nil, err
	}
	return &u, nil
}

// FindByID IDでユーザーを検索
func (c *Client) FindByID(id uuid.UUID) (*user.User, error) {
	var u user.User
	if err := c.db.First(&u, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, pkg.ErrInvalidCredentials
		}
		return nil, err
	}
	return &u, nil
}