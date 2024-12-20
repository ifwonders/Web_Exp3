package takeaway

import (
	"context"
	"gf-demo-takeaway/api/takeaway/v2"
	"gf-demo-takeaway/internal/model"
	"gf-demo-takeaway/internal/service"
)

func (c *ControllerV2) CustomerCreate(ctx context.Context, req *v2.CustomerCreateReq) (res *v2.CustomerCreateRes, err error) {
	res = &v2.CustomerCreateRes{}
	res.One, err = service.Customer().CreateCustomer(ctx, model.CustomerCreateInput{
		CustomerName: req.CustomerName,
		Password:     req.Password,
		Mobile:       req.Mobile,
		Email:        req.Email,
	})
	return
}
