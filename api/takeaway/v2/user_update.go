package v2

import "github.com/gogf/gf/v2/frame/g"

type UserUpdateReq struct {
	g.Meta   `path:"/user/{id}" method:"put" tags:"User" summary:"Update user"`
	Id       int64   `v:"required" dc:"user id"`
	Username *string `v:"required|length:3,10" dc:"username"`
	Password *string `v:"required" dc:"user password"`
	Mobile   *string `dc:"mobile"`
	Email    *string `dc:"email"`
}
type UserUpdateRes struct{}
