package takeaway

import (
	"context"
	"gf-demo-takeaway/internal/dao"
	"gf-demo-takeaway/internal/model/do"

	"gf-demo-takeaway/api/takeaway/v2"
)

func (c *ControllerV2) CustomerGetList(ctx context.Context, req *v2.CustomerGetListReq) (res *v2.CustomerGetListRes, err error) {
	res = &v2.CustomerGetListRes{}
	err = dao.Customer.Ctx(ctx).Where(do.Customer{
		Customername: req.Customername,
		Mobile:       req.Mobile,
		Email:        req.Email,
	}).Scan(&res.List)
	return
}
