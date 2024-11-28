package v2

import "github.com/gogf/gf/v2/frame/g"

type MerchantCreateReq struct {
	g.Meta  `path:"/merchant" method:"post" tags:"merchant" summary:"Create merchant"`
	Name    string `v:"required|length:3,10" dc:"merchant name"`
	Address string `v:"required" dc:"merchant address"`
	Mobile  string `v:"required" dc:"merchant mobile"`
}
type MerchantCreateRes struct {
	Id int64 `json:"id" dc:"merchant id"`
}
