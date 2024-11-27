package takeaway

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"gf-demo-takeaway/api/takeaway/v1"
)

func (c *ControllerV1) TakeawaySelectValue(ctx context.Context, req *v1.TakeawaySelectValueReq) (res *v1.TakeawaySelectValueRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
