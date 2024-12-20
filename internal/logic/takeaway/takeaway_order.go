package takeaway

import (
	"context"
	"gf-demo-takeaway/internal/dao"
	"gf-demo-takeaway/internal/model"
	"gf-demo-takeaway/internal/model/do"
	"gf-demo-takeaway/internal/model/entity"
	"gf-demo-takeaway/internal/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
)

type (
	sOrder struct{}
)

func init() {
	service.RegisterOrder(NewOrder())
}
func NewOrder() service.IOrder {
	return &sOrder{}
}

// CreateOrder creates order account
func (s *sOrder) CreateOrder(ctx context.Context, in model.OrderCreateInput) (record *entity.Order, err error) {
	// 使用事务处理
	var (
		OrderId int64
	)
	ptr := &OrderId
	err = dao.Order.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 执行数据库操作并返回生成的 OrderId
		*ptr, err = dao.Order.Ctx(ctx).Data(do.Order{
			CustomerId:  in.CustomerId,
			MerchantId:  in.MerchantId,
			OrderStatus: in.OrderStatus,
			TotalPrice:  in.TotalPrice,
			CreateTime:  gtime.Now(),
			UpdateTime:  gtime.Now(),
		}).InsertAndGetId()
		// 事务回调函数返回时，先返回 err
		return err
	})
	if err != nil {
		return
	}

	// 事务完成后，返回 err 和 OrderId
	err = dao.Order.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = dao.Order.Ctx(ctx).WherePri(OrderId).Scan(&record)
		return err
	})
	return
}

// DeleteOrder deletes an id-specified order account
func (s *sOrder) DeleteOrder(ctx context.Context, id int64) (err error) {
	return dao.Order.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.Order.Ctx(ctx).WherePri(id).Delete()
		return err
	})
}

// UpdateOrder updates an id-specified order account
func (s *sOrder) UpdateOrder(ctx context.Context, in model.OrderUpdateInput) (record *entity.Order, err error) {
	err = dao.Order.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.Order.Ctx(ctx).Data(do.Order{
			CustomerId:  in.CustomerId,
			MerchantId:  in.MerchantId,
			OrderStatus: in.OrderStatus,
			TotalPrice:  in.TotalPrice,
			UpdateTime:  gtime.Now(),
		}).WherePri(in.Id).Update()
		return err
	})
	if err != nil {
		return
	}
	err = dao.Order.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = dao.Order.Ctx(ctx).WherePri(in.Id).Scan(&record)
		return err
	})
	return
}

// GetOneOrder gets an id-specified account record
func (s *sOrder) GetOneOrder(ctx context.Context, id int64) (Order *entity.Order, err error) {
	err = dao.Order.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = dao.Order.Ctx(ctx).WherePri(id).Scan(&Order)
		return err
	})
	return
}

func (s *sOrder) GetListOrders(ctx context.Context, in model.OrderGetListInput) (Orders []*entity.Order, err error) {
	err = dao.Order.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = dao.Order.Ctx(ctx).Where(do.Order{
			CustomerId: in.CustomerId,
			MerchantId: in.MerchantId,
		}).Scan(&Orders)
		return err
	})
	return
}
