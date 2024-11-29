package takeaway

import (
	"context"
	"gf-demo-takeaway/internal/dao"
	"gf-demo-takeaway/internal/model/do"
	"github.com/gogf/gf/v2/os/gtime"

	"gf-demo-takeaway/api/takeaway/v2"
)

func (c *ControllerV2) CustomerCreate(ctx context.Context, req *v2.CustomerCreateReq) (res *v2.CustomerCreateRes, err error) {
	insertId, err := dao.Customer.Ctx(ctx).Data(do.Customer{
		Customername: req.Customername,
		Password:     req.Password,
		Mobile:       req.Mobile,
		Email:        req.Email,
		CreateTime:   gtime.Now(),
		UpdateTime:   gtime.Now(),
	}).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	res = &v2.CustomerCreateRes{
		Id: insertId,
	}
	return
}
