package payment

import (
	"context"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/payment/model"
)

type PaymentUsecase interface {
	SubmissionWarehouse(ctx context.Context, userId int, req model.SubmissionRequest) (err error)
	GetPaymentScheme(ctx context.Context) (resp []model.SchemeResponse, err error)
}
