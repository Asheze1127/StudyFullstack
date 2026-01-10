package auth

import(
	"backend/internal/authentication/domain/auth"
	"errors"
	"os"
	"strconv"
	"time"
	
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTService struct {
	secretKey string
	tokenExpHours int
}

func NewJWTService() *JWTService {
	secretKey := os.Getenv("JWT_SECRET")

	expHours := 24 // デフォルト24時間
	if expStr := os.Getenv("JWT_EXPIRE_HOURS"); expStr != "" {
		if h, err := strconv.Atoi(expStr); err == nil {
			expHours = h
		}
	}
	
	return &JWTService{
		secretKey: secretKey,
		tokenExpHours: expHours,
	}
}

func (s *JWTService) GenerateToken(userID uuid.UUID, email string) (string, error) {
	claims := &auth.Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(s.tokenExpHours) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}
	
	return tokenString, nil
}

func (s *JWTService) ValidateToken(tokenString string) (*auth.Claims, error) {
	claims := &auth.Claims{}
	
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.secretKey), nil
	})
	
	if err != nil {
		return nil, err
	}
	
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	
	return claims, nil
}
