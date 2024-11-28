package v2

import "github.com/gogf/gf/v2/frame/g"

type OrderStatus int

const (
	OrderStatusPending    OrderStatus = 0 // 订单待处理
	OrderStatusProcessing OrderStatus = 1 // 订单处理中
	OrderStatusShipped    OrderStatus = 2 // 订单已发货
	OrderStatusDelivered  OrderStatus = 3 // 订单已送达
	OrderStatusCancelled  OrderStatus = 4 // 订单已取消
	OrderStatusReturned   OrderStatus = 5 // 订单已退货
)

type OrderCreateReq struct {
	g.Meta      `path:"/order" method:"post" tags:"Order" summary:"Create order"`
	UserId      int         `v:"required" dc:"user id"`
	MerchantId  int         `v:"required" dc:"merchant_id"`
	OrderStatus OrderStatus `v:"required" dc:"order_status"`
	TotalPrice  float32     `v:"required" dc:"total price"`
}
type OrderCreateRes struct {
	Id int64 `json:"id" dc:"order id"`
}
