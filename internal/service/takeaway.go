// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf-demo-takeaway/internal/model"
	"gf-demo-takeaway/internal/model/entity"
)

type (
	ICustomer interface {
		// CreateCustomer creates customer account
		CreateCustomer(ctx context.Context, in model.CustomerCreateInput) (record *entity.Customer, err error)
		// DeleteCustomer deletes an id-specified customer account
		DeleteCustomer(ctx context.Context, id int64) (err error)
		// UpdateCustomer updates an id-specified customer account
		UpdateCustomer(ctx context.Context, in model.CustomerUpdateInput) (record *entity.Customer, err error)
		// GetOneCustomer gets an id-specified account record
		GetOneCustomer(ctx context.Context, id int64) (Customer *entity.Customer, err error)
		GetListCustomers(ctx context.Context, in model.CustomerGetListInput) (Customers []*entity.Customer, err error)
		LoginCustomer(ctx context.Context, in model.CustomerLoginInput) (JWT string, err error)
		ResetPasswordCustomer(ctx context.Context, in model.CustomerResetPasswordInput) (record *entity.Customer, err error)
	}
	IMerchant interface {
		// CreateMerchant creates merchant account
		CreateMerchant(ctx context.Context, in model.MerchantCreateInput) (record *entity.Merchant, err error)
		// DeleteMerchant deletes an id-specified merchant account
		DeleteMerchant(ctx context.Context, id int64) (err error)
		// UpdateMerchant updates an id-specified merchant account
		UpdateMerchant(ctx context.Context, in model.MerchantUpdateInput) (record *entity.Merchant, err error)
		// GetOneMerchant gets an id-specified account record
		GetOneMerchant(ctx context.Context, id int64) (Merchant *entity.Merchant, err error)
		GetListMerchants(ctx context.Context, in model.MerchantGetListInput) (Merchants []*entity.Merchant, err error)
	}
	IOrder interface {
		// CreateOrder creates order account
		CreateOrder(ctx context.Context, in model.OrderCreateInput) (record *entity.Order, err error)
		// DeleteOrder deletes an id-specified order account
		DeleteOrder(ctx context.Context, id int64) (err error)
		// UpdateOrder updates an id-specified order account
		UpdateOrder(ctx context.Context, in model.OrderUpdateInput) (record *entity.Order, err error)
		// GetOneOrder gets an id-specified account record
		GetOneOrder(ctx context.Context, id int64) (Order *entity.Order, err error)
		GetListOrders(ctx context.Context, in model.OrderGetListInput) (Orders []*entity.Order, err error)
	}
)

var (
	localCustomer ICustomer
	localMerchant IMerchant
	localOrder    IOrder
)

func Customer() ICustomer {
	if localCustomer == nil {
		panic("implement not found for interface ICustomer, forgot register?")
	}
	return localCustomer
}

func RegisterCustomer(i ICustomer) {
	localCustomer = i
}

func Merchant() IMerchant {
	if localMerchant == nil {
		panic("implement not found for interface IMerchant, forgot register?")
	}
	return localMerchant
}

func RegisterMerchant(i IMerchant) {
	localMerchant = i
}

func Order() IOrder {
	if localOrder == nil {
		panic("implement not found for interface IOrder, forgot register?")
	}
	return localOrder
}

func RegisterOrder(i IOrder) {
	localOrder = i
}
