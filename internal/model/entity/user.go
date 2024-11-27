// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id         int         `json:"id"         orm:"id"          description:""` //
	Username   string      `json:"username"   orm:"username"    description:""` //
	Password   string      `json:"password"   orm:"password"    description:""` //
	Mobile     string      `json:"mobile"     orm:"mobile"      description:""` //
	Email      string      `json:"email"      orm:"email"       description:""` //
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:""` //
	UpdateTime *gtime.Time `json:"updateTime" orm:"update_time" description:""` //
}
