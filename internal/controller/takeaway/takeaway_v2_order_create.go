package takeaway

import (
	"context"
	"gf-demo-takeaway/internal/dao"
	"gf-demo-takeaway/internal/model/do"
	"github.com/gogf/gf/v2/os/gtime"

	"gf-demo-takeaway/api/takeaway/v2"
)

func (c *ControllerV2) OrderCreate(ctx context.Context, req *v2.OrderCreateReq) (res *v2.OrderCreateRes, err error) {
	insertId, err := dao.Order.Ctx(ctx).Data(do.Order{
		UserId:      req.UserId,
		MerchantId:  req.MerchantId,
		OrderStatus: req.OrderStatus,
		TotalPrice:  req.TotalPrice,
		CreateTime:  gtime.Now(),
		UpdateTime:  gtime.Now(),
	}).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	res = &v2.OrderCreateRes{
		Id: insertId,
	}
	return
}
