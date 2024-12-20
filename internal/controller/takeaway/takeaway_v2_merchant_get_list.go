package takeaway

import (
	"context"
	"gf-demo-takeaway/internal/model"
	"gf-demo-takeaway/internal/service"

	"gf-demo-takeaway/api/takeaway/v2"
)

func (c *ControllerV2) MerchantGetList(ctx context.Context, req *v2.MerchantGetListReq) (res *v2.MerchantGetListRes, err error) {
	res = &v2.MerchantGetListRes{}
	res.List, err = service.Merchant().GetListMerchants(ctx, model.MerchantGetListInput{
		Name: req.Name,
	})
	return
}
