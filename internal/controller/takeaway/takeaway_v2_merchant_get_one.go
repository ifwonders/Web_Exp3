package takeaway

import (
	"context"
	"gf-demo-takeaway/internal/service"

	"gf-demo-takeaway/api/takeaway/v2"
)

func (c *ControllerV2) MerchantGetOne(ctx context.Context, req *v2.MerchantGetOneReq) (res *v2.MerchantGetOneRes, err error) {
	res = &v2.MerchantGetOneRes{}
	res.One, err = service.Merchant().GetOneMerchant(ctx, req.Id)
	return
}
