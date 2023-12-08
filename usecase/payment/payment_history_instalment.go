package payment

import (
	"context"
	"fmt"
	"net/http"
	"time"

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
		var paymentTime time.Time
		if list[i].OngoingInstalment != nil {
			paymentTime = *list[i].OngoingInstalment.PaymentTime
		}
		data := model.TransactionHistoryResponse{
			TransactionID:     list[i].TransactionID,
			InstalmentId:      list[i].ID,
			TransactionDate:   paymentTime,
			PaymentSchemeId:   list[i].Transaction.PaymentScheme.ID,
			PaymentSchemeName: list[i].Transaction.PaymentScheme.Scheme,
			UserID:            list[i].Transaction.User.ID,
			UserName:          list[i].Transaction.User.Username,
			Nominal:           list[i].Nominal,
		}

		resp = append(resp, data)
	}
	return
}
