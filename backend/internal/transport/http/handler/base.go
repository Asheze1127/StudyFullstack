package handler

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

// BaseHandler ベースハンドラー（共通処理）
type BaseHandler struct{}

func NewBaseHandler() *BaseHandler {
	return &BaseHandler{}
}

// RespondJSON JSONレスポンスを返す
func (h *BaseHandler) RespondJSON(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}

// RespondError エラーレスポンスを返す
func (h *BaseHandler) RespondError(c *gin.Context, status int, err error) {
	c.JSON(status, gin.H{
		"error": err.Error(),
	})
}

// RespondValidationError バリデーションエラーレスポンスを返す
func (h *BaseHandler) RespondValidationError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
}
