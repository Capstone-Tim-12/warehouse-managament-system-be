package paymentdb

import (
	"context"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/entity"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/paginate"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	FindPaymentSchemeById(ctx context.Context, id int) (resp *entity.PaymentScheme, err error)
	CreateTransaction(ctx context.Context, req *entity.Transaction) (err error)
	FindPaymentScheme(ctx context.Context) (resp []entity.PaymentScheme, err error)
	GetListTransactionDasboar(ctx context.Context, param paginate.Pagination) (resp []entity.Transaction, count int64, err error)
	GetTransactionByUserId(ctx context.Context, userId int) (resp []entity.Transaction, err error)
	GetInstalmentUser(ctx context.Context, param paginate.Pagination) (resp []entity.Instalment, count int64, err error)
	GetListTransactionData(ctx context.Context, param paginate.PaginationTrx) (resp []entity.Transaction, count int64, err error)
	GetTransactionById(ctx context.Context, transactionId string) (resp *entity.Transaction, err error)
	BeginTrans(ctx context.Context) *gorm.DB
	CreateInstalment(ctx context.Context, tx *gorm.DB, req *entity.Instalment) (err error)
	UpdateTransaction(ctx context.Context, tx *gorm.DB, req *entity.Transaction) (err error)
}
