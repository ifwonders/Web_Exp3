package takeaway

import (
	"context"
	"gf-demo-takeaway/internal/dao"

	"gf-demo-takeaway/api/takeaway/v2"
)

func (c *ControllerV2) UserGetOne(ctx context.Context, req *v2.UserGetOneReq) (res *v2.UserGetOneRes, err error) {
	res = &v2.UserGetOneRes{}
	err = dao.User.Ctx(ctx).WherePri(req.Id).Scan(&res.User)
	return
}
