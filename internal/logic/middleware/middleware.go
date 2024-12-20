package middleware

import (
	"gf-demo-takeaway/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v4"
)

type (
	sMiddleware struct{}
)

func init() {
	service.RegisterMiddleware(New())
}

func New() service.IMiddleware {
	return &sMiddleware{}
}

// UserClaims 用于存储用户信息的 JWT Claims
type UserClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// AuthMiddleware 认证中间件，用于验证 JWT Token
func (s *sMiddleware) AuthMiddleware(r *ghttp.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		r.Response.WriteJsonExit(g.Map{"code": 401, "msg": "Authorization token required"})
	}

	// 解析和验证 JWT Token
	token, err := service.JWT().ParseJWT(authHeader)
	if err != nil || !token.Valid {
		r.Response.WriteJsonExit(g.Map{"code": 401, "msg": "Invalid or expired token"})
	}

	// 获取 Claims 信息
	//claims, ok := token.Claims.(*UserClaims)
	//if !ok {
	//	r.Response.WriteJsonExit(g.Map{"code": 401, "msg": "Invalid token"})
	//}

	//// 将用户信息存入上下文
	//r.SetCtxVar("username", claims.Username)
	r.Middleware.Next()
}
