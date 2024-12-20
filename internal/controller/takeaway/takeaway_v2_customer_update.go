package takeaway

import (
	"context"
	"gf-demo-takeaway/api/takeaway/v2"
	"gf-demo-takeaway/internal/model"
	"gf-demo-takeaway/internal/service"
)

func (c *ControllerV2) CustomerUpdate(ctx context.Context, req *v2.CustomerUpdateReq) (res *v2.CustomerUpdateRes, err error) {
	res = &v2.CustomerUpdateRes{}
	res.One, err = service.Customer().UpdateCustomer(ctx, model.CustomerUpdateInput{
		Id:           req.Id,
		CustomerName: *req.CustomerName,
		Password:     *req.Password,
		Mobile:       *req.Mobile,
		Email:        *req.Email,
	})
	return
}
