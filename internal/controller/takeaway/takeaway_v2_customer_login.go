package takeaway

import (
	"context"
	"gf-demo-takeaway/api/takeaway/v2"
	"gf-demo-takeaway/internal/model"
	"gf-demo-takeaway/internal/service"
)

func (c *ControllerV2) CustomerLogin(ctx context.Context, req *v2.CustomerLoginReq) (res *v2.CustomerLoginRes, err error) {
	JWT, err := service.Customer().LoginCustomer(ctx, model.CustomerLoginInput{
		CustomerName: req.CustomerName,
		Password:     req.Password,
	})
	if err != nil {
		return
	}
	res = &v2.CustomerLoginRes{
		Token: JWT,
	}
	return
}
