package takeaway

import (
	"context"
	"gf-demo-takeaway/internal/dao"

	"gf-demo-takeaway/api/takeaway/v2"
)

func (c *ControllerV2) CustomerDelete(ctx context.Context, req *v2.CustomerDeleteReq) (res *v2.CustomerDeleteRes, err error) {
	_, err = dao.Customer.Ctx(ctx).WherePri(req.Id).Delete()
	return
}
