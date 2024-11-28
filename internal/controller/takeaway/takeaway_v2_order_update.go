package takeaway

import (
	"context"
	"gf-demo-takeaway/internal/dao"
	"gf-demo-takeaway/internal/model/do"
	"github.com/gogf/gf/v2/os/gtime"

	"gf-demo-takeaway/api/takeaway/v2"
)

func (c *ControllerV2) OrderUpdate(ctx context.Context, req *v2.OrderUpdateReq) (res *v2.OrderUpdateRes, err error) {
	_, err = dao.Order.Ctx(ctx).Data(do.Order{
		UserId:      req.UserId,
		MerchantId:  req.MerchantId,
		OrderStatus: req.OrderStatus,
		TotalPrice:  req.TotalPrice,
		UpdateTime:  gtime.Now(),
	}).WherePri(req.Id).Update()
	return
}
