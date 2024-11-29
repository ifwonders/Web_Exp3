package v2

import "github.com/gogf/gf/v2/frame/g"

type CustomerCreateReq struct {
	g.Meta       `path:"/customer" method:"post" tags:"Customer" summary:"Create customer"`
	Customername string `v:"required|length:3,10" dc:"customername"`
	Password     string `v:"required" dc:"customer password"`
	Mobile       string `dc:"mobile"`
	Email        string `dc:"email"`
}
type CustomerCreateRes struct {
	Id int64 `json:"id" dc:"customer id"`
}
