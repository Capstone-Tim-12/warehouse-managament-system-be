package payment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/payment/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
)

func (s *defaultPayment) GetListTransactionIdUser(ctx context.Context, userId int) (resp []model.ListTrxUser, err error) {
	trxData, err := s.paymentRepo.GetTransactionByUserId(ctx, userId)
	if err != nil {
		fmt.Println("error getting transaction: ", err.Error())
		err = errors.New(http.StatusInternalServerError, "error getting transaction")
		return
	}

	for i := 0; i < len(trxData); i++ {
		resp = append(resp, model.ListTrxUser{
			TransactionId: trxData[i].ID,
		})
	}
	return
}