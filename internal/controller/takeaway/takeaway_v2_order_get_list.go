package takeaway

import (
	"context"
	"gf-demo-takeaway/internal/model"
	"gf-demo-takeaway/internal/service"

	"gf-demo-takeaway/api/takeaway/v2"
)

func (c *ControllerV2) OrderGetList(ctx context.Context, req *v2.OrderGetListReq) (res *v2.OrderGetListRes, err error) {
	res = &v2.OrderGetListRes{}
	res.List, err = service.Order().GetListOrders(ctx, model.OrderGetListInput{
		CustomerId: req.CustomerId,
		MerchantId: req.MerchantId,
	})
	return
}
