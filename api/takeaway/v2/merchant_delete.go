package v2

import "github.com/gogf/gf/v2/frame/g"

type MerchantDeleteReq struct {
	g.Meta `path:"/merchant/{id}" method:"delete" tags:"Merchant" summary:"Delete merchant"`
	Id     int64 `v:"required" dc:"merchant id"`
}
type MerchantDeleteRes struct{}
