package v2

import "github.com/gogf/gf/v2/frame/g"

type MerchantUpdateReq struct {
	g.Meta  `path:"/merchant/{id}" method:"put" tags:"Merchant" summary:"Update merchant"`
	Id      int64   `v:"required" dc:"merchant id"`
	Name    *string `v:"required|length:3,10" dc:"merchant name"`
	Address *string `v:"required" dc:"merchant address"`
	Mobile  *string `v:"required" dc:"merchant mobile"`
}
type MerchantUpdateRes struct{}
