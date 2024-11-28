package v2

import "github.com/gogf/gf/v2/frame/g"

type UserCreateReq struct {
	g.Meta   `path:"/user" method:"post" tags:"User" summary:"Create user"`
	Username string `v:"required|length:3,10" dc:"username"`
	Password string `v:"required" dc:"user password"`
	Mobile   string `dc:"mobile"`
	Email    string `dc:"email"`
}
type UserCreateRes struct {
	Id int64 `json:"id" dc:"user id"`
}
