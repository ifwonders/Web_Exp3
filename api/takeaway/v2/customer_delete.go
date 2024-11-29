package v2

import "github.com/gogf/gf/v2/frame/g"

type CustomerDeleteReq struct {
	g.Meta `path:"/customer/{id}" method:"delete" tags:"Customer" summary:"Delete customer"`
	Id     int64 `v:"required" dc:"customer id"`
}
type CustomerDeleteRes struct{}
