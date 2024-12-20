// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package takeaway

import (
	"context"

	"gf-demo-takeaway/api/takeaway/v2"
)

type ITakeawayV2 interface {
	CustomerCreate(ctx context.Context, req *v2.CustomerCreateReq) (res *v2.CustomerCreateRes, err error)
	CustomerDelete(ctx context.Context, req *v2.CustomerDeleteReq) (res *v2.CustomerDeleteRes, err error)
	CustomerLogin(ctx context.Context, req *v2.CustomerLoginReq) (res *v2.CustomerLoginRes, err error)
	CustomerLogout(ctx context.Context, req *v2.CustomerLogoutReq) (res *v2.CustomerLogoutRes, err error)
	CustomerResetPassword(ctx context.Context, req *v2.CustomerResetPasswordReq) (res *v2.CustomerResetPasswordRes, err error)
	CustomerGetOne(ctx context.Context, req *v2.CustomerGetOneReq) (res *v2.CustomerGetOneRes, err error)
	CustomerGetList(ctx context.Context, req *v2.CustomerGetListReq) (res *v2.CustomerGetListRes, err error)
	CustomerUpdate(ctx context.Context, req *v2.CustomerUpdateReq) (res *v2.CustomerUpdateRes, err error)
	EmailSendAuthCode(ctx context.Context, req *v2.EmailSendAuthCodeReq) (res *v2.EmailSendAuthCodeRes, err error)
	EmailVerifyAuthCode(ctx context.Context, req *v2.EmailVerifyAuthCodeReq) (res *v2.EmailVerifyAuthCodeRes, err error)
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
}
