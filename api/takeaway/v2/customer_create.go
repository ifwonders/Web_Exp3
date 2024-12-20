package v2

import (
	"gf-demo-takeaway/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type CustomerCreateReq struct {
	g.Meta       `path:"/customer" method:"post" tags:"Customer" summary:"Create customer"`
	CustomerName string `v:"required|length:3,10" dc:"customer name"`
	Password     string `v:"required" dc:"customer password"`
	Mobile       string `dc:"mobile"`
	Email        string `dc:"email"`
}
type CustomerCreateRes struct {
	One *entity.Customer `dc:"customer"`
}
