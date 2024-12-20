package v2

import "github.com/gogf/gf/v2/frame/g"

type CustomerLoginReq struct {
	g.Meta       `path:"/customer/login" method:"post" tags:"Customer" summary:"login customer"`
	CustomerName string `json:"customer_name"`
	Password     string `json:"password"`
}
type CustomerLoginRes struct {
	Token string `json:"token"`
}
