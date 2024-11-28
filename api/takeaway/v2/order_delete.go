package v2

import "github.com/gogf/gf/v2/frame/g"

type OrderDeleteReq struct {
	g.Meta `path:"/order/{id}" method:"delete" tags:"Order" summary:"Delete order"`
	Id     int64 `v:"required" dc:"order id"`
}
type OrderDeleteRes struct{}
