// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Authcode is the golang structure of table authcode for DAO operations like Where/Data.
type Authcode struct {
	g.Meta      `orm:"table:authcode, do:true"`
	Id          interface{} //
	Email       interface{} //
	Authcode    interface{} //
	ExpiredTime *gtime.Time //
}
