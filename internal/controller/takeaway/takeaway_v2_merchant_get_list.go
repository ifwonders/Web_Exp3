package takeaway

import (
	"context"
	"gf-demo-takeaway/internal/dao"
	"gf-demo-takeaway/internal/model/do"

	"gf-demo-takeaway/api/takeaway/v2"
)

func (c *ControllerV2) MerchantGetList(ctx context.Context, req *v2.MerchantGetListReq) (res *v2.MerchantGetListRes, err error) {
	res = &v2.MerchantGetListRes{}
	err = dao.Merchant.Ctx(ctx).Where(do.Merchant{
		Name:    req.Name,
		Mobile:  req.Mobile,
		Address: req.Address,
	}).Scan(&res.List)
	return
}
