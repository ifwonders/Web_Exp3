package takeaway

import (
	"context"
	"gf-demo-takeaway/api/takeaway/v2"
	"gf-demo-takeaway/internal/model"
	"gf-demo-takeaway/internal/service"
)

func (c *ControllerV2) OrderUpdate(ctx context.Context, req *v2.OrderUpdateReq) (res *v2.OrderUpdateRes, err error) {
	res = &v2.OrderUpdateRes{}
	res.One, err = service.Order().UpdateOrder(ctx, model.OrderUpdateInput{
		Id:          req.Id,
		CustomerId:  *req.CustomerId,
		MerchantId:  *req.MerchantId,
		OrderStatus: *req.OrderStatus,
		TotalPrice:  *req.TotalPrice,
	})
	return
}
