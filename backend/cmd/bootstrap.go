package main

import (
	authhandler "backend/internal/authentication/handler"
	authservice "backend/internal/authentication/service/auth"
	userservice "backend/internal/authentication/service/user"
	"backend/internal/authentication/usecase"
	"backend/internal/transport/http/handler"
	"backend/router"
	"errors"
	"fmt"
	"os"
	
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Bootstrap 依存性注入と初期化
func Bootstrap() (*gin.Engine, error) {
	// データベース接続
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// DATABASE_URLが設定されていない場合、個別の環境変数をチェック
		host := os.Getenv("DB_HOST")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME")
		port := os.Getenv("DB_PORT")
		sslmode := os.Getenv("DB_SSLMODE")
		
		// 必須環境変数のチェック
		if host == "" {
			return nil, errors.New("DB_HOST environment variable is required")
		}
		if user == "" {
			return nil, errors.New("DB_USER environment variable is required")
		}
		if password == "" {
			return nil, errors.New("DB_PASSWORD environment variable is required")
		}
		if dbname == "" {
			return nil, errors.New("DB_NAME environment variable is required")
		}
		if port == "" {
			return nil, errors.New("DB_PORT environment variable is required")
		}
		if sslmode == "" {
			return nil, errors.New("DB_SSLMODE environment variable is required")
		}
		
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			host, user, password, dbname, port, sslmode)
	}
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	
	// サービス実装
	userSvc := userservice.NewClient(db)
	authSvc := authservice.NewClient()
	
	// ユースケース
	signInUC := usecase.NewSignInUseCase(userSvc, authSvc)
	signUpUC := usecase.NewSignUpUseCase(userSvc, authSvc)
	
	// ハンドラー
	baseHandler := handler.NewBaseHandler()
	authHandler := authhandler.NewAuthHandler(baseHandler, signInUC, signUpUC)
	
	// ルーター設定
	r := router.SetupRouter(authHandler)
	
	return r, nil
}