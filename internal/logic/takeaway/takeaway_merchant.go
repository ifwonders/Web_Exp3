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
	sMerchant struct{}
)

func init() {
	service.RegisterMerchant(NewMerchant())
}
func NewMerchant() service.IMerchant {
	return &sMerchant{}
}

// CreateMerchant creates merchant account
func (s *sMerchant) CreateMerchant(ctx context.Context, in model.MerchantCreateInput) (record *entity.Merchant, err error) {
	var (
		MerchantId int64
	)
	ptr := &MerchantId
	// 使用事务处理
	err = dao.Merchant.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 执行数据库操作并返回生成的 MerchantId
		*ptr, err = dao.Merchant.Ctx(ctx).Data(do.Merchant{
			Name:       in.Name,
			Address:    in.Address,
			Mobile:     in.Mobile,
			CreateTime: gtime.Now(),
			UpdateTime: gtime.Now(),
		}).InsertAndGetId()
		// 事务回调函数返回时，先返回 err
		return err
	})
	if err != nil {
		return
	}
	// 事务完成后，返回 err 和 MerchantId
	err = dao.Merchant.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = dao.Merchant.Ctx(ctx).WherePri(MerchantId).Scan(&record)
		return err
	})
	return
}

// DeleteMerchant deletes an id-specified merchant account
func (s *sMerchant) DeleteMerchant(ctx context.Context, id int64) (err error) {
	return dao.Merchant.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.Merchant.Ctx(ctx).WherePri(id).Delete()
		return err
	})
}

// UpdateMerchant updates an id-specified merchant account
func (s *sMerchant) UpdateMerchant(ctx context.Context, in model.MerchantUpdateInput) (record *entity.Merchant, err error) {
	err = dao.Merchant.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.Merchant.Ctx(ctx).Data(do.Merchant{
			Name:       in.Name,
			Address:    in.Address,
			Mobile:     in.Mobile,
			UpdateTime: gtime.Now(),
		}).WherePri(in.Id).Update()
		return err
	})
	if err != nil {

		return
	}
	err = dao.Merchant.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = dao.Merchant.Ctx(ctx).WherePri(in.Id).Scan(&record)
		return err
	})
	return
}

// GetOneMerchant gets an id-specified account record
func (s *sMerchant) GetOneMerchant(ctx context.Context, id int64) (Merchant *entity.Merchant, err error) {
	err = dao.Merchant.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = dao.Merchant.Ctx(ctx).WherePri(id).Scan(&Merchant)
		return err
	})
	return
}

func (s *sMerchant) GetListMerchants(ctx context.Context, in model.MerchantGetListInput) (Merchants []*entity.Merchant, err error) {
	err = dao.Merchant.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = dao.Merchant.Ctx(ctx).Where(do.Merchant{
			Name: in.Name,
		}).Scan(&Merchants)
		return err
	})
	return
}
