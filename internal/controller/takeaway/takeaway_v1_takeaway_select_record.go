package takeaway

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"gf-demo-takeaway/api/takeaway/v1"
)

func (c *ControllerV1) TakeawaySelectRecord(ctx context.Context, req *v1.TakeawaySelectRecordReq) (res *v1.TakeawaySelectRecordRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
