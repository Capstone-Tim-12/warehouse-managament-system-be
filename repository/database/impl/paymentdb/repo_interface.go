package paymentdb

import (
	"context"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/entity"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/paginate"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	FindPaymentSchemeById(ctx context.Context, id int) (resp *entity.PaymentScheme, err error)
	CreateTransaction(ctx context.Context, tx *gorm.DB, req *entity.Transaction) (err error)
	FindPaymentScheme(ctx context.Context) (resp []entity.PaymentScheme, err error)
	GetListTransactionDasboar(ctx context.Context, param paginate.Pagination) (resp []entity.Transaction, count int64, err error)
	GetTransactionByUserId(ctx context.Context, userId int, param paginate.Pagination) (resp []entity.Transaction, count int64, err error)
	GetInstalmentUser(ctx context.Context, param paginate.Pagination) (resp []entity.Instalment, count int64, err error)
	GetListTransactionData(ctx context.Context, param paginate.PaginationTrx) (resp []entity.Transaction, count int64, err error)
	GetTransactionById(ctx context.Context, transactionId string) (resp *entity.Transaction, err error)
	BeginTrans(ctx context.Context) *gorm.DB
	CreateInstalment(ctx context.Context, tx *gorm.DB, req *entity.Instalment) (err error)
	UpdateTransaction(ctx context.Context, tx *gorm.DB, req *entity.Transaction) (err error)
	GetTransactionDetailById(ctx context.Context, transactionId string) (resp *entity.Transaction, err error)
	GetListTransactionByUserIdAndStatus(ctx context.Context, userId int, status entity.TranscationStatus, param paginate.Pagination) (resp []entity.Transaction, count int64, err error)
	GetListInstalmentByTransactionId(ctx context.Context, transactionId string, param paginate.Pagination) (resp []entity.Instalment, count int64, err error)
	GetListPaymentMethod(ctx context.Context) (resp []entity.PaymentMethod, err error)
	GetPaymentMethodById(ctx context.Context, id int) (resp *entity.PaymentMethod, err error)
	GetInstalmentById(ctx context.Context, id int) (resp *entity.Instalment, err error)
	UpdateInstalment(ctx context.Context, tx *gorm.DB, req *entity.Instalment) (err error)
	CreateOngoingInstalment(ctx context.Context, tx *gorm.DB, req *entity.OngoingInstalment) (err error)
	UpdateOngoingInstalment(ctx context.Context, tx *gorm.DB, req *entity.OngoingInstalment) (err error)
	FindOngoingInstalmentByXpayment(ctx context.Context, xpaymentId string) (resp *entity.OngoingInstalment, err error)
	GetTransactionUserDetailByTransactionId(ctx context.Context, transactionId string) (resp *entity.Transaction, err error)
	GetTransactionDetailByWarehouseId(ctx context.Context, warehouseId int, param paginate.Pagination) (resp []entity.Transaction, count int64, err error)
	GetTotalPayment(ctx context.Context) (totalPayment float64, err error)
	GetTotalPaymentOnYear(ctx context.Context, year int) (totalPayment float64, err error)
}
