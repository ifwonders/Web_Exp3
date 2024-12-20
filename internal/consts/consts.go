package consts

type OrderStatus int

const (
	OrderStatusPending    OrderStatus = 0 // 订单待处理
	OrderStatusProcessing OrderStatus = 1 // 订单处理中
	OrderStatusShipped    OrderStatus = 2 // 订单已发货
	OrderStatusDelivered  OrderStatus = 3 // 订单已送达
	OrderStatusCancelled  OrderStatus = 4 // 订单已取消
	OrderStatusReturned   OrderStatus = 5 // 订单已退货
)
