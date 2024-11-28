package takeaway

import (
	"context"
	"gf-demo-takeaway/internal/dao"

	"gf-demo-takeaway/api/takeaway/v2"
)

func (c *ControllerV2) MerchantGetOne(ctx context.Context, req *v2.MerchantGetOneReq) (res *v2.MerchantGetOneRes, err error) {
	res = &v2.MerchantGetOneRes{}
	err = dao.Merchant.Ctx(ctx).WherePri(req.Id).Scan(&res.Merchant)
	return
}
