package paymentdb

import (
	"context"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/entity"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/paginate"
	"gorm.io/gorm"
)

type defaultRepo struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &defaultRepo{db: db}
}

func (s *defaultRepo) FindPaymentSchemeById(ctx context.Context, id int) (resp *entity.PaymentScheme, err error) {
	err = s.db.WithContext(ctx).Take(&resp, "id = ?", id).Error
	return
}

func (s *defaultRepo) CreateTransaction(ctx context.Context, req *entity.Transaction) (err error) {
	err = s.db.WithContext(ctx).Create(req).Error
	return
}

func (s *defaultRepo) FindPaymentScheme(ctx context.Context) (resp []entity.PaymentScheme, err error) {
	err = s.db.WithContext(ctx).Find(&resp).Error
	return
}

func (s *defaultRepo) GetListTransactionDasboar(ctx context.Context, param paginate.Pagination) (resp []entity.Transaction, count int64, err error) {
	err = s.db.WithContext(ctx).Model(&entity.Transaction{}).Count(&count).Error
	if err != nil {
		return
	}
	err = s.db.WithContext(ctx).
	Preload("User").Scopes(paginate.Paginate(param.Page, param.Limit)).Find(&resp).Error
	return
}

func (s *defaultRepo) GetInstalmentUser(ctx context.Context, param paginate.Pagination) (resp []entity.Instalment, count int64, err error) {
	query := func (condition *gorm.DB) *gorm.DB  {
		condition.Where("status = ? or status = ? ", entity.Paid, entity.Failed)
		return condition
	}

	err = s.db.WithContext(ctx).Model(&entity.Instalment{}).Scopes(query).Count(&count).Error
	if err != nil {
		return
	}
	err = s.db.WithContext(ctx).Scopes(paginate.Paginate(param.Page, param.Limit)).
		Scopes(query).Preload("Transaction").Preload("Transaction.User").Find(&resp).Error
	return
}

func (s *defaultRepo) GetTransactionByUserId(ctx context.Context, userId int) (resp []entity.Transaction, err error) {
	err  = s.db.WithContext(ctx).Find(&resp, "user_id = ?", userId).Error
	return
}
