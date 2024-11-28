package v2

import "github.com/gogf/gf/v2/frame/g"

type UserDeleteReq struct {
	g.Meta `path:"/user/{id}" method:"delete" tags:"User" summary:"Delete user"`
	Id     int64 `v:"required" dc:"user id"`
}
type UserDeleteRes struct{}
