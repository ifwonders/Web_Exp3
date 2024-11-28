package takeaway

import (
	"context"
	"gf-demo-takeaway/internal/dao"

	"gf-demo-takeaway/api/takeaway/v2"
)

func (c *ControllerV2) MerchantDelete(ctx context.Context, req *v2.MerchantDeleteReq) (res *v2.MerchantDeleteRes, err error) {
	_, err = dao.Merchant.Ctx(ctx).WherePri(req.Id).Delete()
	return
}
