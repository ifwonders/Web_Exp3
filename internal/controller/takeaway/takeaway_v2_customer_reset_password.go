package takeaway

import (
	"context"
	"gf-demo-takeaway/internal/model"
	"gf-demo-takeaway/internal/service"

	"gf-demo-takeaway/api/takeaway/v2"
)

func (c *ControllerV2) CustomerResetPassword(ctx context.Context, req *v2.CustomerResetPasswordReq) (res *v2.CustomerResetPasswordRes, err error) {
	res = &v2.CustomerResetPasswordRes{}
	res.One, err = service.Customer().ResetPasswordCustomer(ctx, model.CustomerResetPasswordInput{
		Id:       req.Id,
		Password: req.Password,
	})
	return
}
