// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Customer is the golang structure of table customer for DAO operations like Where/Data.
type Customer struct {
	g.Meta       `orm:"table:customer, do:true"`
	Id           interface{} //
	Customername interface{} //
	Password     interface{} //
	Mobile       interface{} //
	Email        interface{} //
	CreateTime   *gtime.Time //
	UpdateTime   *gtime.Time //
}
