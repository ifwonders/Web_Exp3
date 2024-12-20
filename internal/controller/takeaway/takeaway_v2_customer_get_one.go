package takeaway

import (
	"context"
	"gf-demo-takeaway/api/takeaway/v2"
	"gf-demo-takeaway/internal/service"
)

func (c *ControllerV2) CustomerGetOne(ctx context.Context, req *v2.CustomerGetOneReq) (res *v2.CustomerGetOneRes, err error) {
	res = &v2.CustomerGetOneRes{}
	res.One, err = service.Customer().GetOneCustomer(ctx, req.Id)
	return
}
