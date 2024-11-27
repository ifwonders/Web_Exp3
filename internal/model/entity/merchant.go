// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Merchant is the golang structure for table merchant.
type Merchant struct {
	Id         int         `json:"id"         orm:"id"          description:""` //
	Name       string      `json:"name"       orm:"name"        description:""` //
	Address    string      `json:"address"    orm:"address"     description:""` //
	Mobile     string      `json:"mobile"     orm:"mobile"      description:""` //
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:""` //
	UpdateTime *gtime.Time `json:"updateTime" orm:"update_time" description:""` //
}
