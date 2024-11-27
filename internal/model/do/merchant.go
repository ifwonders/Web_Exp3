// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Merchant is the golang structure of table merchant for DAO operations like Where/Data.
type Merchant struct {
	g.Meta     `orm:"table:merchant, do:true"`
	Id         interface{} //
	Name       interface{} //
	Address    interface{} //
	Mobile     interface{} //
	CreateTime *gtime.Time //
	UpdateTime *gtime.Time //
}
