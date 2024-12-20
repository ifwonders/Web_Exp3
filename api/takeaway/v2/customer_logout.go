package v2

import "github.com/gogf/gf/v2/frame/g"

type CustomerLogoutReq struct {
	g.Meta `path:"/protected/customer_logout" method:"get" tags:"Customer" summary:"logout customers"`
}

type CustomerLogoutRes struct{}
