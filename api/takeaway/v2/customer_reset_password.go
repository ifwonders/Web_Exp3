package v2

import (
	"gf-demo-takeaway/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type CustomerResetPasswordReq struct {
	g.Meta   `path:"/protected/reset_password" method:"post" tags:"Customer" summary:"Reset Password"`
	Id       int64  `v:"required" dc:"customer id"`
	Password string `json:"password"`
}
type CustomerResetPasswordRes struct {
	One *entity.Customer `dc:"customer"`
}
