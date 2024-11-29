package v2

import (
	"gf-demo-takeaway/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type CustomerGetOneReq struct {
	g.Meta `path:"/customer/{id}" method:"get" tags:"Customer" summary:"Get one customer"`
	Id     int64 `v:"required" dc:"customer id"`
}
type CustomerGetOneRes struct {
	*entity.Customer `dc:"customer"`
}
type CustomerGetListReq struct {
	g.Meta       `path:"/customer" method:"get" tags:"Customer" summary:"Get customers"`
	Customername string `v:"length:3,10" dc:"customername"`
	Mobile       string `dc:"mobile"`
	Email        string `dc:"email"`
}
type CustomerGetListRes struct {
	List []*entity.Customer `json:"list" dc:"customer list"`
}
