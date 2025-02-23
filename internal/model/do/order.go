// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Order is the golang structure of table order for DAO operations like Where/Data.
type Order struct {
	g.Meta      `orm:"table:order, do:true"`
	Id          interface{} //
	CustomerId  interface{} //
	MerchantId  interface{} //
	OrderStatus interface{} //
	TotalPrice  interface{} //
	CreateTime  *gtime.Time //
	UpdateTime  *gtime.Time //
}
