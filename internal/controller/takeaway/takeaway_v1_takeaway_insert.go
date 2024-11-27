package takeaway

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"gf-demo-takeaway/api/takeaway/v1"
)

func (c *ControllerV1) TakeawayInsert(ctx context.Context, req *v1.TakeawayInsertReq) (res *v1.TakeawayInsertRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
