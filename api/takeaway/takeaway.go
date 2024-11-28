// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package takeaway

import (
	"context"

	"gf-demo-takeaway/api/takeaway/v2"
)

type ITakeawayV2 interface {
	MerchantCreate(ctx context.Context, req *v2.MerchantCreateReq) (res *v2.MerchantCreateRes, err error)
	MerchantDelete(ctx context.Context, req *v2.MerchantDeleteReq) (res *v2.MerchantDeleteRes, err error)
	MerchantGetOne(ctx context.Context, req *v2.MerchantGetOneReq) (res *v2.MerchantGetOneRes, err error)
	MerchantGetList(ctx context.Context, req *v2.MerchantGetListReq) (res *v2.MerchantGetListRes, err error)
	MerchantUpdate(ctx context.Context, req *v2.MerchantUpdateReq) (res *v2.MerchantUpdateRes, err error)
	OrderCreate(ctx context.Context, req *v2.OrderCreateReq) (res *v2.OrderCreateRes, err error)
	OrderDelete(ctx context.Context, req *v2.OrderDeleteReq) (res *v2.OrderDeleteRes, err error)
	OrderGetOne(ctx context.Context, req *v2.OrderGetOneReq) (res *v2.OrderGetOneRes, err error)
	OrderGetList(ctx context.Context, req *v2.OrderGetListReq) (res *v2.OrderGetListRes, err error)
	OrderUpdate(ctx context.Context, req *v2.OrderUpdateReq) (res *v2.OrderUpdateRes, err error)
	UserCreate(ctx context.Context, req *v2.UserCreateReq) (res *v2.UserCreateRes, err error)
	UserDelete(ctx context.Context, req *v2.UserDeleteReq) (res *v2.UserDeleteRes, err error)
	UserGetOne(ctx context.Context, req *v2.UserGetOneReq) (res *v2.UserGetOneRes, err error)
	UserGetList(ctx context.Context, req *v2.UserGetListReq) (res *v2.UserGetListRes, err error)
	UserUpdate(ctx context.Context, req *v2.UserUpdateReq) (res *v2.UserUpdateRes, err error)
}
