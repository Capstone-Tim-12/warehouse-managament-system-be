package paymentdb

import (
	"context"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/entity"
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

func (s *defaultRepo) GetListTransactionDasboar(ctx context.Context) (resp []entity.Transaction, err error) {
	err = s.db.WithContext(ctx).Preload("User").Find(&resp).Error
	return
}

func (s *defaultRepo) GetTransactionByUserId(ctx context.Context, userId int) (resp []entity.Transaction, err error) {
	err  = s.db.WithContext(ctx).Find(&resp, "user_id = ?", userId).Error
	return
}
