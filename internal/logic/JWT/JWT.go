package JWT

import (
	"fmt"
	"gf-demo-takeaway/internal/service"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type (
	sJWT struct{}
)

func init() {
	service.RegisterJWT(NewJWT())
}

func NewJWT() service.IJWT {
	return &(sJWT{})
}

var jwtSecret = []byte("jwt_secret_key") // 密钥应保密，实际生产中请使用环境变量管理

// UserClaims 用于存储用户信息的 JWT Claims
type UserClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateJWT 生成 JWT Token
func (s *sJWT) GenerateJWT(username string) (string, error) {
	claims := UserClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "gf-demo-takeaway",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)), // 1 小时过期
		},
	}

	// 使用 HMAC SHA256 签名方法生成 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ParseJWT 解析 JWT Token
func (s *sJWT) ParseJWT(tokenString string) (*jwt.Token, error) {
	// 解析 token 时，验证签名和有效性
	return jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 确保 token 使用的是 HMAC SHA256 签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
}
