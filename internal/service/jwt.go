// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"github.com/golang-jwt/jwt/v4"
)

type (
	IJWT interface {
		// GenerateJWT 生成 JWT Token
		GenerateJWT(username string) (string, error)
		// ParseJWT 解析 JWT Token
		ParseJWT(tokenString string) (*jwt.Token, error)
	}
)

var (
	localJWT IJWT
)

func JWT() IJWT {
	if localJWT == nil {
		panic("implement not found for interface IJWT, forgot register?")
	}
	return localJWT
}

func RegisterJWT(i IJWT) {
	localJWT = i
}
