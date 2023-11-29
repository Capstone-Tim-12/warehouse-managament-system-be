package paymentdb

import (
	"context"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/entity"
)

type PaymentRepository interface {
	FindPaymentSchemeById(ctx context.Context, id int) (resp *entity.PaymentScheme, err error)
	CreateTransaction(ctx context.Context, req *entity.Transaction) (err error)
	FindPaymentScheme(ctx context.Context) (resp []entity.PaymentScheme, err error)
}
