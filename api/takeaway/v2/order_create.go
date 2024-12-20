package v2

import (
	"gf-demo-takeaway/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)
import "gf-demo-takeaway/internal/consts"

type OrderCreateReq struct {
	g.Meta      `path:"/order" method:"post" tags:"Order" summary:"Create order"`
	CustomerId  int64              `v:"required" dc:"customer id"`
	MerchantId  int64              `v:"required" dc:"merchant_id"`
	OrderStatus consts.OrderStatus `v:"required" dc:"order_status"`
	TotalPrice  float64            `v:"required" dc:"total price"`
}
type OrderCreateRes struct {
	One *entity.Order `dc:"order"`
}
