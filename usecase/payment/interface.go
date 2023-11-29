package payment

import (
	"context"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/payment/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/paginate"
)

type PaymentUsecase interface {
	SubmissionWarehouse(ctx context.Context, userId int, req model.SubmissionRequest) (err error)
	GetPaymentScheme(ctx context.Context, id int) (resp []model.SchemeResponse, err error)
	HistoryTransactions(ctx context.Context, param paginate.Pagination) (resp []model.TransactionHistoryResponse, count int64, err error)
}
