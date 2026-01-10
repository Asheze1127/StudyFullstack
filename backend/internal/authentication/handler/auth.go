package handler

import (
	"backend/internal/authentication/usecase"
	"backend/internal/transport/http/handler"
	"backend/pkg"
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// AuthHandler 認証ハンドラー
type AuthHandler struct {
	*handler.BaseHandler
	signInUC usecase.SignInUseCase
	signUpUC usecase.SignUpUseCase
}

// NewAuthHandler コンストラクタ
func NewAuthHandler(
	base *handler.BaseHandler,
	signInUC usecase.SignInUseCase,
	signUpUC usecase.SignUpUseCase,
) *AuthHandler {
	return &AuthHandler{
		BaseHandler: base,
		signInUC:    signInUC,
		signUpUC:    signUpUC,
	}
}

// SignInRequest サインインリクエスト
type SignInRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

// SignInResponse サインインレスポンス
type SignInResponse struct {
	User struct {
		ID    uuid.UUID `json:"id"`    // uint → uuid.UUID に変更
		Email string    `json:"email"`
	} `json:"user"`
	Token string `json:"token"`
}

// SignIn godoc
// @Summary      ユーザーログイン
// @Description  メールアドレスとパスワードでログイン
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      SignInRequest  true  "ログイン情報"
// @Success      200      {object}  SignInResponse
// @Failure      400      {object}  ErrorResponse
// @Failure      401      {object}  ErrorResponse
// @Router       /v1/users/auth/login [post]
func (h *AuthHandler) SignIn(c *gin.Context) {
	var req SignInRequest
	
	// バリデーション
	if err := c.ShouldBindJSON(&req); err != nil {
		h.RespondValidationError(c, err)
		return
	}
	
	// ユースケース実行
	params := &usecase.SignInParams{
		Email:    req.Email,
		Password: req.Password,
	}
	
	output, err := h.signInUC.Execute(params)
	if err != nil {
		status := http.StatusUnauthorized
		if err == pkg.ErrInvalidCredentials {
			status = http.StatusUnauthorized
		} else {
			status = http.StatusInternalServerError
		}
		h.RespondError(c, status, err)
		return
	}
	
	// レスポンス作成
	var res SignInResponse
	res.User.ID = output.UserID
	res.User.Email = output.Email
	res.Token = output.Token
	
	h.RespondJSON(c, http.StatusOK, res)
}

// SignUpRequest サインアップリクエスト
type SignUpRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

// SignUpResponse サインアップレスポンス
type SignUpResponse struct {
	User struct {
		ID    uuid.UUID `json:"id"`    // uint → uuid.UUID に変更
		Email string    `json:"email"`
	} `json:"user"`
}

// SignUp godoc
// @Summary      ユーザー登録
// @Description  新規ユーザーを登録
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      SignUpRequest  true  "登録情報"
// @Success      201      {object}  SignUpResponse
// @Failure      400      {object}  ErrorResponse
// @Failure      409      {object}  ErrorResponse
// @Router       /v1/users/auth/signup [post]
func (h *AuthHandler) SignUp(c *gin.Context) {
	var req SignUpRequest
	
	// バリデーション
	if err := c.ShouldBindJSON(&req); err != nil {
		h.RespondValidationError(c, err)
		return
	}
	
	// ユースケース実行
	params := &usecase.SignUpParams{
		Email:    req.Email,
		Password: req.Password,
	}
	
	output, err := h.signUpUC.Execute(params)
	if err != nil {
		status := http.StatusInternalServerError
		if err == pkg.ErrEmailAlreadyExists {
			status = http.StatusConflict
		}
		h.RespondError(c, status, err)
		return
	}
	
	// レスポンス作成
	var res SignUpResponse
	res.User.ID = output.UserID
	res.User.Email = output.Email
	
	h.RespondJSON(c, http.StatusCreated, res)
}

// ErrorResponse エラーレスポンス
type ErrorResponse struct {
	Error string `json:"error"`
}