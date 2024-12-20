package takeaway

import (
	"context"
	"gf-demo-takeaway/internal/service"

	"gf-demo-takeaway/api/takeaway/v2"
)

func (c *ControllerV2) OrderGetOne(ctx context.Context, req *v2.OrderGetOneReq) (res *v2.OrderGetOneRes, err error) {
	res = &v2.OrderGetOneRes{}
	res.One, err = service.Order().GetOneOrder(ctx, req.Id)
	return
}
