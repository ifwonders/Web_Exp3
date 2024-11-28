package takeaway

import (
	"context"
	"gf-demo-takeaway/internal/dao"
	"gf-demo-takeaway/internal/model/do"
	"github.com/gogf/gf/v2/os/gtime"

	"gf-demo-takeaway/api/takeaway/v2"
)

func (c *ControllerV2) UserUpdate(ctx context.Context, req *v2.UserUpdateReq) (res *v2.UserUpdateRes, err error) {
	_, err = dao.User.Ctx(ctx).Data(do.User{
		Username:   req.Username,
		Password:   req.Password,
		Mobile:     req.Mobile,
		Email:      req.Email,
		UpdateTime: gtime.Now(),
	}).WherePri(req.Id).Update()
	return
}
