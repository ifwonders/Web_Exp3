package takeaway

import (
	"context"
	"gf-demo-takeaway/internal/dao"
	"gf-demo-takeaway/internal/model/do"
	"github.com/gogf/gf/v2/os/gtime"

	"gf-demo-takeaway/api/takeaway/v2"
)

func (c *ControllerV2) MerchantCreate(ctx context.Context, req *v2.MerchantCreateReq) (res *v2.MerchantCreateRes, err error) {
	insertId, err := dao.Merchant.Ctx(ctx).Data(do.Merchant{
		Name:       req.Name,
		Address:    req.Address,
		Mobile:     req.Mobile,
		CreateTime: gtime.Now(),
		UpdateTime: gtime.Now(),
	}).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	res = &v2.MerchantCreateRes{
		Id: insertId,
	}
	return
}
