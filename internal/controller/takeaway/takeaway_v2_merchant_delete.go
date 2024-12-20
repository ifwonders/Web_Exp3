package takeaway

import (
	"context"
	"gf-demo-takeaway/internal/service"

	"gf-demo-takeaway/api/takeaway/v2"
)

func (c *ControllerV2) MerchantDelete(ctx context.Context, req *v2.MerchantDeleteReq) (res *v2.MerchantDeleteRes, err error) {
	err = service.Merchant().DeleteMerchant(ctx, req.Id)
	return
}
