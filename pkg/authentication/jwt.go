package authentication

import (
	"fmt"
	"time"

	"chitchat4.0/pkg/model"
	"github.com/golang-jwt/jwt/v4"
)

const (
	Issuer = "hpj.io"
)

// CustomClaims 定制要求
type CustomClaims struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

// JWTService 结构体
type JWTService struct {
	signKey        []byte
	issuer         string
	expireDuration time.Duration
}

// NewJWTService 创建一个 JWT 服务
func NewJWTService(secret string) *JWTService {
	return &JWTService{
		signKey:        []byte(secret),
		issuer:         Issuer,
		expireDuration: 7 * 24 * time.Hour,
	}
}

// ParseToken 解析token
func (s *JWTService) ParseToken(tokenString string) (*model.User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return s.signKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*CustomClaims)
	if !ok || token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	user := &model.User{
		ID:   claims.ID,
		Name: claims.Name,
	}
	return user, nil
}
