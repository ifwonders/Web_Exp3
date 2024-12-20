package takeaway

import (
	"context"
	"gf-demo-takeaway/api/takeaway/v2"
	"gf-demo-takeaway/internal/model"
	"gf-demo-takeaway/internal/service"
)

func (c *ControllerV2) CustomerGetList(ctx context.Context, req *v2.CustomerGetListReq) (res *v2.CustomerGetListRes, err error) {
	res = &v2.CustomerGetListRes{}
	res.List, err = service.Customer().GetListCustomers(ctx, model.CustomerGetListInput{
		CustomerName: req.CustomerName,
	})
	return
}
