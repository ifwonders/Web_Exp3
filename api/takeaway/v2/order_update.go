package v2

import (
	"gf-demo-takeaway/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)
import "gf-demo-takeaway/internal/consts"

type OrderUpdateReq struct {
	g.Meta      `path:"/order/{id}" method:"put" tags:"Order" summary:"Update order"`
	Id          int64               `v:"required" dc:"order id"`
	CustomerId  *int64              `v:"required" dc:"customer id"`
	MerchantId  *int64              `v:"required" dc:"merchant id"`
	OrderStatus *consts.OrderStatus `v:"required" dc:"order status"`
	TotalPrice  *float64            `v:"required" dc:"total price"`
}
type OrderUpdateRes struct {
	One *entity.Order `dc:"order"`
}
