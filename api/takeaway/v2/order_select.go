package v2

import (
	"gf-demo-takeaway/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type OrderGetOneReq struct {
	g.Meta `path:"/order/{id}" method:"get" tags:"Order" summary:"Get one order"`
	Id     int64 `v:"required" dc:"order id"`
}
type OrderGetOneRes struct {
	*entity.Order `dc:"order"`
}
type OrderGetListReq struct {
	g.Meta     `path:"/order" method:"get" tags:"Order" summary:"Get orders"`
	CustomerId int `dc:"customer id"`
	MerchantId int `dc:"merchant id"`
}
type OrderGetListRes struct {
	List []*entity.Order `json:"list" dc:"order list"`
}
