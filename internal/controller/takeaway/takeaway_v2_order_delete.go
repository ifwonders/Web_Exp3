package takeaway

import (
	"context"
	"gf-demo-takeaway/internal/dao"

	"gf-demo-takeaway/api/takeaway/v2"
)

func (c *ControllerV2) OrderDelete(ctx context.Context, req *v2.OrderDeleteReq) (res *v2.OrderDeleteRes, err error) {
	_, err = dao.Order.Ctx(ctx).WherePri(req.Id).Delete()
	return
}
