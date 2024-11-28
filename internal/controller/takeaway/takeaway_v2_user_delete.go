package takeaway

import (
	"context"
	"gf-demo-takeaway/internal/dao"

	"gf-demo-takeaway/api/takeaway/v2"
)

func (c *ControllerV2) UserDelete(ctx context.Context, req *v2.UserDeleteReq) (res *v2.UserDeleteRes, err error) {
	_, err = dao.User.Ctx(ctx).WherePri(req.Id).Delete()
	return
}
