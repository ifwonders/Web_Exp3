package takeaway

import (
	"context"
	"gf-demo-takeaway/api/takeaway/v2"
	"gf-demo-takeaway/internal/model"
	"gf-demo-takeaway/internal/service"
)

func (c *ControllerV2) MerchantCreate(ctx context.Context, req *v2.MerchantCreateReq) (res *v2.MerchantCreateRes, err error) {
	res = &v2.MerchantCreateRes{}
	res.One, err = service.Merchant().CreateMerchant(ctx, model.MerchantCreateInput{
		Name:    req.Name,
		Address: req.Address,
		Mobile:  req.Mobile,
	})
	return
}
