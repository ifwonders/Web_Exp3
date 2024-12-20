package takeaway

import (
	"context"
	"gf-demo-takeaway/api/takeaway/v2"
	"gf-demo-takeaway/internal/model"
	"gf-demo-takeaway/internal/service"
)

func (c *ControllerV2) MerchantUpdate(ctx context.Context, req *v2.MerchantUpdateReq) (res *v2.MerchantUpdateRes, err error) {
	res = &v2.MerchantUpdateRes{}
	res.One, err = service.Merchant().UpdateMerchant(ctx, model.MerchantUpdateInput{
		Id:      req.Id,
		Name:    *req.Name,
		Address: *req.Address,
		Mobile:  *req.Mobile,
	})
	return
}
