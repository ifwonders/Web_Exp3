package takeaway

import (
	"context"
	"gf-demo-takeaway/internal/dao"
	"gf-demo-takeaway/internal/model/do"

	"gf-demo-takeaway/api/takeaway/v2"
)

func (c *ControllerV2) UserGetList(ctx context.Context, req *v2.UserGetListReq) (res *v2.UserGetListRes, err error) {
	res = &v2.UserGetListRes{}
	err = dao.User.Ctx(ctx).Where(do.User{
		Username: req.Username,
		Mobile:   req.Mobile,
		Email:    req.Email,
	}).Scan(&res.List)
	return
}
