package takeaway

import (
	"context"
	"gf-demo-takeaway/internal/service"

	"gf-demo-takeaway/api/takeaway/v2"
)

func (c *ControllerV2) OrderDelete(ctx context.Context, req *v2.OrderDeleteReq) (res *v2.OrderDeleteRes, err error) {
	err = service.Order().DeleteOrder(ctx, req.Id)
	return
}
