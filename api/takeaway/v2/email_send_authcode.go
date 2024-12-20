package v2

import "github.com/gogf/gf/v2/frame/g"

type EmailSendAuthCodeReq struct {
	g.Meta `path:"/email/authcode/send" method:"post" tags:"Email" summary:"Get email authcode"`
	Email  string `v:"required"`
}
type EmailSendAuthCodeRes struct {
	Id int64 `json:"id"`
}
