package payment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/payment/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
)

func (s *defaultPayment) GetTransactionInfo(ctx context.Context, transactionId string) (resp model.TransactionInfoResponse, err error) {
	trxData, err := s.paymentRepo.GetTransactionById(ctx, transactionId)
	if err != nil {
		fmt.Println("error getting transaction")
		err = errors.New(http.StatusNotFound, "transaction not found")
		return
	}

	resp = model.TransactionInfoResponse{
		DateEntry:     trxData.DateEntry,
		DateOut:       trxData.DateOut,
		PaymentScheme: trxData.PaymentScheme.Scheme,
		Duration:      trxData.Duration,
		Username:      trxData.User.Username,
		Email:         trxData.User.Email,
	}

	return
}
