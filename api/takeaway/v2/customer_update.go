package v2

import "github.com/gogf/gf/v2/frame/g"

type CustomerUpdateReq struct {
	g.Meta       `path:"/customer/{id}" method:"put" tags:"Customer" summary:"Update customer"`
	Id           int64   `v:"required" dc:"customer id"`
	Customername *string `v:"required|length:3,10" dc:"customername"`
	Password     *string `v:"required" dc:"customer password"`
	Mobile       *string `dc:"mobile"`
	Email        *string `dc:"email"`
}
type CustomerUpdateRes struct{}
