package router

import (
	authhandler "backend/internal/authentication/handler"
	
	"github.com/gin-gonic/gin"
)

// SetupRouter ルーター設定
func SetupRouter(authHandler *authhandler.AuthHandler) *gin.Engine {
	r := gin.Default()
	
	// API v1
	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/users/auth")
		{
			authGroup.POST("/login", authHandler.SignIn)
			authGroup.POST("/signup", authHandler.SignUp)
		}
	}
	
	return r
}