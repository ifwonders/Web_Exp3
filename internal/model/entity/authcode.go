// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Authcode is the golang structure for table authcode.
type Authcode struct {
	Id          int         `json:"id"          orm:"id"           description:""` //
	Email       string      `json:"email"       orm:"email"        description:""` //
	Authcode    string      `json:"authcode"    orm:"authcode"     description:""` //
	ExpiredTime *gtime.Time `json:"expiredTime" orm:"expired_time" description:""` //
}
