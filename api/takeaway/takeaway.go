// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package takeaway

import (
	"context"

	"gf-demo-takeaway/api/takeaway/v1"
)

type ITakeawayV1 interface {
	TakeawayDelete(ctx context.Context, req *v1.TakeawayDeleteReq) (res *v1.TakeawayDeleteRes, err error)
	TakeawayInsert(ctx context.Context, req *v1.TakeawayInsertReq) (res *v1.TakeawayInsertRes, err error)
	TakeawaySelectValue(ctx context.Context, req *v1.TakeawaySelectValueReq) (res *v1.TakeawaySelectValueRes, err error)
	TakeawaySelectRecord(ctx context.Context, req *v1.TakeawaySelectRecordReq) (res *v1.TakeawaySelectRecordRes, err error)
	TakeawayUpdate(ctx context.Context, req *v1.TakeawayUpdateReq) (res *v1.TakeawayUpdateRes, err error)
}
