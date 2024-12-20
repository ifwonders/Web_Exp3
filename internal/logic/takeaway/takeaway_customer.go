package takeaway

import (
	"context"
	"gf-demo-takeaway/internal/dao"
	"gf-demo-takeaway/internal/model"
	"gf-demo-takeaway/internal/model/do"
	"gf-demo-takeaway/internal/model/entity"
	"gf-demo-takeaway/internal/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

type (
	sCustomer struct{}
)

func init() {
	service.RegisterCustomer(NewCustomer())
}
func NewCustomer() service.ICustomer {
	return &sCustomer{}
}

// CreateCustomer creates customer account
func (s *sCustomer) CreateCustomer(ctx context.Context, in model.CustomerCreateInput) (record *entity.Customer, err error) {
	var (
		CustomerId int64
	)
	ptr := &CustomerId
	// 使用事务处理
	err = dao.Customer.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 执行数据库操作并返回生成的 CustomerId
		*ptr, err = dao.Customer.Ctx(ctx).Data(do.Customer{
			Customername: in.CustomerName,
			Password:     in.Password,
			Mobile:       in.Mobile,
			Email:        in.Email,
			CreateTime:   gtime.Now(),
			UpdateTime:   gtime.Now(),
		}).InsertAndGetId()
		// 事务回调函数返回时，先返回 err
		return err
	})
	err = dao.Customer.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = dao.Customer.Ctx(ctx).WherePri(*ptr).Scan(&record)
		return err
	})
	return
}

// DeleteCustomer deletes an id-specified customer account
func (s *sCustomer) DeleteCustomer(ctx context.Context, id int64) (err error) {
	return dao.Customer.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.Customer.Ctx(ctx).WherePri(id).Delete()
		return err
	})
}

// UpdateCustomer updates an id-specified customer account
func (s *sCustomer) UpdateCustomer(ctx context.Context, in model.CustomerUpdateInput) (record *entity.Customer, err error) {
	err = dao.Customer.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.Customer.Ctx(ctx).Data(do.Customer{
			Customername: in.CustomerName,
			Password:     in.Password,
			Mobile:       in.Mobile,
			Email:        in.Email,
			UpdateTime:   gtime.Now(),
		}).WherePri(in.Id).Update()
		return err
	})
	err = dao.Customer.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = dao.Customer.Ctx(ctx).WherePri(in.Id).Scan(&record)
		return err
	})
	return
}

// GetOneCustomer gets an id-specified account record
func (s *sCustomer) GetOneCustomer(ctx context.Context, id int64) (Customer *entity.Customer, err error) {
	err = dao.Customer.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = dao.Customer.Ctx(ctx).WherePri(id).Scan(&Customer)
		return err
	})
	return
}

func (s *sCustomer) GetListCustomers(ctx context.Context, in model.CustomerGetListInput) (Customers []*entity.Customer, err error) {
	err = dao.Customer.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = dao.Customer.Ctx(ctx).Where(do.Customer{
			Customername: in.CustomerName,
		}).Scan(&Customers)
		return err
	})
	return
}

func (s *sCustomer) LoginCustomer(ctx context.Context, in model.CustomerLoginInput) (JWT string, err error) {
	var customer *entity.Customer
	err = dao.Customer.Ctx(ctx).Where(do.Customer{
		Customername: in.CustomerName,
		Password:     in.Password,
	}).Scan(&customer)
	if err != nil {
		return "ERROR", err
	}
	if customer == nil {
		return "ERROR", gerror.New(`Passport or Password not correct`)
	}
	JWT, err = service.JWT().GenerateJWT(customer.CustomerName)
	return
}

func (s *sCustomer) ResetPasswordCustomer(ctx context.Context, in model.CustomerResetPasswordInput) (record *entity.Customer, err error) {
	err = dao.Customer.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.Customer.Ctx(ctx).Data(do.Customer{
			Password:   in.Password,
			UpdateTime: gtime.Now(),
		}).WherePri(in.Id).Update()
		return err
	})
	if err != nil {
		return
	}
	err = dao.Customer.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = dao.Customer.Ctx(ctx).WherePri(in.Id).Scan(&record)
		return err
	})
	return
}
