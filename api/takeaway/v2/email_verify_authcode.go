package v2

import (
	"github.com/gogf/gf/v2/frame/g"
)

type EmailVerifyAuthCodeReq struct {
	g.Meta   `path:"/email/authcode/verify" method:"post" tags:"Email" summary:"Verify email authcode"`
	Email    string `v:"required"`
	Authcode string `v:"required"`
	//VerifyTime gtime.Time `v:""`
}
type EmailVerifyAuthCodeRes struct {
	VerifyResult bool `v:"-"`
}
