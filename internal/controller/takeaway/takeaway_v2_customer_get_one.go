package takeaway

import (
	"context"
	"gf-demo-takeaway/internal/dao"

	"gf-demo-takeaway/api/takeaway/v2"
)

func (c *ControllerV2) CustomerGetOne(ctx context.Context, req *v2.CustomerGetOneReq) (res *v2.CustomerGetOneRes, err error) {
	res = &v2.CustomerGetOneRes{}
	err = dao.Customer.Ctx(ctx).WherePri(req.Id).Scan(&res.Customer)
	return
}
