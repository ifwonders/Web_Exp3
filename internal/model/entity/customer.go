// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Customer is the golang structure for table customer.
type Customer struct {
	Id           int         `json:"id"         orm:"id"          description:""`         //
	Customername string      `json:"customername"   orm:"customername"    description:""` //
	Password     string      `json:"password"   orm:"password"    description:""`         //
	Mobile       string      `json:"mobile"     orm:"mobile"      description:""`         //
	Email        string      `json:"email"      orm:"email"       description:""`         //
	CreateTime   *gtime.Time `json:"createTime" orm:"create_time" description:""`         //
	UpdateTime   *gtime.Time `json:"updateTime" orm:"update_time" description:""`         //
}
