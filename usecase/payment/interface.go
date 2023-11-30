package payment

import (
	"context"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/payment/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/paginate"
)

type PaymentUsecase interface {
	SubmissionWarehouse(ctx context.Context, userId int, req model.SubmissionRequest) (err error)
	GetPaymentScheme(ctx context.Context) (resp []model.SchemeResponse, err error)
	GetHistoryInstalmentUser(ctx context.Context, param paginate.Pagination) (resp []model.TransactionHistoryResponse, count int64, err error)
	GetListTransactionByUserId(ctx context.Context, userId int) (resp []model.ListTrxUser, err error)
	GetAllTransaction(ctx context.Context, param paginate.PaginationTrx) (resp []model.ListAllTrxResponse, count int64, err error) 
	TransactionApproved(ctx context.Context, transactionId string) (err error)
	TransactionRejected(ctx context.Context, transactionId string) (err error)
}
