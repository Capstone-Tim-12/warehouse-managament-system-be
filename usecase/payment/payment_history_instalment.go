package payment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/payment/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/paginate"
)

func (s *defaultPayment) GetHistoryInstalmentUser(ctx context.Context, param paginate.Pagination) (resp []model.TransactionHistoryResponse, count int64, err error) {
	list, count, err := s.paymentRepo.GetInstalmentUser(ctx, param)
	if err != nil {
		fmt.Println("error getting transactionList: ", err.Error())
		err = errors.New(http.StatusInternalServerError, "error getting transaction list")
		return
	}

	for i := 0; i < len(list); i++ {
		resp = append(resp, model.TransactionHistoryResponse{
			TransactionID:     list[i].TransactionID,
			InstalmentId:      list[i].ID,
			UserID:            list[i].Transaction.User.ID,
			UserName:          list[i].Transaction.User.Username,
			StatusTransaction: string(list[i].Status),
			Nominal:           list[i].Nominal,
		})
	}
	return
}