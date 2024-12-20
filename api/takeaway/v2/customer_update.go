package v2

import (
	"gf-demo-takeaway/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type CustomerUpdateReq struct {
	g.Meta       `path:"/protected/customer/{id}" method:"put" tags:"Customer" summary:"Update customer"`
	Id           int64   `v:"required" dc:"customer id"`
	CustomerName *string `v:"required|length:3,10" dc:"customer name"`
	Password     *string `v:"required" dc:"customer password"`
	Mobile       *string `dc:"mobile"`
	Email        *string `dc:"email"`
}
type CustomerUpdateRes struct {
	One *entity.Customer `dc:"customer"`
}
