package v2

import (
	"gf-demo-takeaway/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type MerchantGetOneReq struct {
	g.Meta `path:"/merchant/{id}" method:"get" tags:"Merchant" summary:"Get one merchant"`
	Id     int64 `v:"required" dc:"merchant id"`
}
type MerchantGetOneRes struct {
	One *entity.Merchant `dc:"merchant"`
}
type MerchantGetListReq struct {
	g.Meta  `path:"/merchant" method:"get" tags:"Merchant" summary:"Get merchants"`
	Name    string `v:"length:3,10" dc:"merchant name"`
	Address string `dc:"merchant address"`
	Mobile  string `dc:"merchant mobile"`
}
type MerchantGetListRes struct {
	List []*entity.Merchant `json:"list" dc:"merchant list"`
}
