package takeaway

import (
	"context"
	"gf-demo-takeaway/api/takeaway/v2"
	"gf-demo-takeaway/internal/model"
	"gf-demo-takeaway/internal/service"
)

func (c *ControllerV2) OrderCreate(ctx context.Context, req *v2.OrderCreateReq) (res *v2.OrderCreateRes, err error) {
	res = &v2.OrderCreateRes{}
	res.One, err = service.Order().CreateOrder(ctx, model.OrderCreateInput{
		CustomerId:  req.CustomerId,
		MerchantId:  req.MerchantId,
		OrderStatus: req.OrderStatus,
		TotalPrice:  req.TotalPrice,
	})
	return
}
