package v2

import "github.com/gogf/gf/v2/frame/g"

type OrderUpdateReq struct {
	g.Meta      `path:"/order/{id}" method:"put" tags:"order" summary:"Update order"`
	Id          int64        `v:"required" dc:"order id"`
	UserId      *int         `v:"required" dc:"user id"`
	MerchantId  *int         `v:"required" dc:"merchant id"`
	OrderStatus *OrderStatus `v:"required" dc:"order status"`
	TotalPrice  *float32     `v:"required" dc:"total price"`
}
type OrderUpdateRes struct{}
