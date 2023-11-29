package payment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/payment/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/paginate"
)

func (s *defaultPayment) HistoryTransactions(ctx context.Context, param paginate.Pagination) (resp []model.TransactionHistoryResponse, count int64, err error) {
	list, count, err := s.paymentRepo.GetListTransactionDasboar(ctx, param)
	if err != nil {
		fmt.Println("error getting transactionList: ", err.Error())
		err = errors.New(http.StatusInternalServerError, "error getting transaction list")
		return
	}

	for i := 0; i < len(list); i++ {
		userData, _ := s.userRepo.GetUserById(ctx, list[i].UserID)
		var nominal float64
		if len(list[i].Instalment) != 0 {
			nominal = list[i].Instalment[0].Nominal
		}

		resp = append(resp, model.TransactionHistoryResponse{
			TransactionID:     list[i].ID,
			UserID:            list[i].UserID,
			UserName:          userData.Username,
			StatusTransaction: string(list[i].Status),
			Nominal:           nominal,
		})
	}
	return
}
