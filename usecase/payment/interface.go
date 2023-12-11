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
	GetAllTransaction(ctx context.Context, param paginate.PaginationTrx) (resp []model.ListAllTrxResponse, count int64, err error) 
	TransactionApproved(ctx context.Context, transactionId string) (err error)
	TransactionRejected(ctx context.Context, transactionId string) (err error)
	GetTransactionListDetail(ctx context.Context, transactionId string) (resp model.TrxListDetail, err error) 
	GetListInstalmentByTrxId(ctx context.Context, transactionId string, param paginate.Pagination) (resp []model.ListInstalmentResponse, count int64, err error)
	GetTransactionInfo(ctx context.Context, transactionId string) (resp model.TransactionInfoResponse, err error)
	GetListPaymentMethod(ctx context.Context) (resp []model.PaymentMethodResponse, err error)
	GetBankVa(ctx context.Context) (resp []model.VaBankResponse, err error) 
	PaymentCheckout(ctx context.Context, userId int, req model.PaymentRequest) (resp model.PaymentResponse, err error)
	VaCallback(ctx context.Context, req model.VaCallbackRequest) (err error)
	GetListTrxUserDasboar(ctx context.Context, userId int, param paginate.Pagination) (resp []model.ListTrxUserDasboarResponse, count int64, err error)
	GetTransactionDetailDasboardUser(ctx context.Context, transactionId string) (resp model.TransactionDetailUser, err error)
	GetListTranscationByWarehouseId(ctx context.Context, warehouseId int, param paginate.Pagination) (resp []model.ListTransactionWarehouseDasboard, count int64, err error)
	GetTotalPaymentDasboard(ctx context.Context) (resp model.GetTotalPayment, err error)
	GetStatistictPaymentOnYear(ctx context.Context) (resp []model.StatiscticPayment, err error)
	GetReasoneList(ctx context.Context) (resp []model.ReasoneData, err error)
}
