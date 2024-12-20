package takeaway

import (
	"context"
	"gf-demo-takeaway/api/takeaway/v2"
	"gf-demo-takeaway/internal/service"
)

func (c *ControllerV2) CustomerDelete(ctx context.Context, req *v2.CustomerDeleteReq) (res *v2.CustomerDeleteRes, err error) {
	err = service.Customer().DeleteCustomer(ctx, req.Id)
	return
}
