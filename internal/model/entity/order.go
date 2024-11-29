// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Order is the golang structure for table order.
type Order struct {
	Id          int         `json:"id"          orm:"id"           description:""`         //
	CustomerId  int         `json:"customerId"      orm:"customer_id"      description:""` //
	MerchantId  int         `json:"merchantId"  orm:"merchant_id"  description:""`         //
	OrderStatus int         `json:"orderStatus" orm:"order_status" description:""`         //
	TotalPrice  float64     `json:"totalPrice"  orm:"total_price"  description:""`         //
	CreateTime  *gtime.Time `json:"createTime"  orm:"create_time"  description:""`         //
	UpdateTime  *gtime.Time `json:"updateTime"  orm:"update_time"  description:""`         //
}
