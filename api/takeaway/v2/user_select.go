package v2

import (
	"gf-demo-takeaway/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type UserGetOneReq struct {
	g.Meta `path:"/user/{id}" method:"get" tags:"User" summary:"Get one user"`
	Id     int64 `v:"required" dc:"user id"`
}
type UserGetOneRes struct {
	*entity.User `dc:"user"`
}
type UserGetListReq struct {
	g.Meta   `path:"/user" method:"get" tags:"User" summary:"Get users"`
	Username string `v:"length:3,10" dc:"username"`
	Mobile   string `dc:"mobile"`
	Email    string `dc:"email"`
}
type UserGetListRes struct {
	List []*entity.User `json:"list" dc:"user list"`
}
