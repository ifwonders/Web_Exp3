package takeaway

import (
	"context"

	"gf-demo-takeaway/api/takeaway/v2"
)

func (c *ControllerV2) CustomerLogout(ctx context.Context, req *v2.CustomerLogoutReq) (res *v2.CustomerLogoutRes, err error) {
	// 注销只需删除客户端的 JWT Token，服务器端不需要进行状态更新
	return
}
