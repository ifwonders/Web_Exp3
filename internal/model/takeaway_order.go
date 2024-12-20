package model

import "gf-demo-takeaway/internal/consts"

type OrderCreateInput struct {
	CustomerId  int64              `json:"customer_id"`
	MerchantId  int64              `json:"merchant_id"`
	OrderStatus consts.OrderStatus `json:"order_status"`
	TotalPrice  float64            `json:"total_price"`
}

type OrderUpdateInput struct {
	Id          int64              `json:"id"`
	CustomerId  int64              `json:"customer_id"`
	MerchantId  int64              `json:"merchant_id"`
	OrderStatus consts.OrderStatus `json:"order_status"`
	TotalPrice  float64            `json:"total_price"`
}

type OrderGetListInput struct {
	CustomerId int64 `json:"customer_id"`
	MerchantId int64 `json:"merchant_id"`
}
