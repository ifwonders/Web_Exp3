package takeaway

import (
	"context"
	"gf-demo-takeaway/internal/dao"
	"gf-demo-takeaway/internal/model/do"
	"github.com/gogf/gf/v2/os/gtime"

	"gf-demo-takeaway/api/takeaway/v2"
)

func (c *ControllerV2) UserCreate(ctx context.Context, req *v2.UserCreateReq) (res *v2.UserCreateRes, err error) {
	insertId, err := dao.User.Ctx(ctx).Data(do.User{
		Username:   req.Username,
		Password:   req.Password,
		Mobile:     req.Mobile,
		Email:      req.Email,
		CreateTime: gtime.Now(),
		UpdateTime: gtime.Now(),
	}).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	res = &v2.UserCreateRes{
		Id: insertId,
	}
	return
}
