package takeaway

import (
	"context"
	"gf-demo-takeaway/internal/dao"
	"gf-demo-takeaway/internal/model/do"

	"gf-demo-takeaway/api/takeaway/v2"
)

func (c *ControllerV2) OrderGetList(ctx context.Context, req *v2.OrderGetListReq) (res *v2.OrderGetListRes, err error) {
	res = &v2.OrderGetListRes{}
	err = dao.Order.Ctx(ctx).Where(do.Order{
		CustomerId: req.CustomerId,
		MerchantId: req.MerchantId,
	}).Scan(&res.List)
	return
}
