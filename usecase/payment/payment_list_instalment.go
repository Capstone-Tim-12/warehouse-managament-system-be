package payment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/payment/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/paginate"
)

func (s *defaultPayment) GetListInstalmentByTrxId(ctx context.Context, transactionId string, param paginate.Pagination) (resp []model.ListInstalmentResponse, count int64, err error) {
	trxData, count, err := s.paymentRepo.GetListInstalmentByTransactionId(ctx, transactionId, param)
	if err != nil {
		fmt.Println("error getting list instalment: ", err.Error())
		err = errors.New(http.StatusInternalServerError, "error getting list instalment")
		return
	}
	for i := 0; i < len(trxData); i++ {
		resp = append(resp, model.ListInstalmentResponse{
			InstalmentId: trxData[i].ID,
			DueDate:      trxData[i].DueDate,
			Nominal:      int(trxData[i].Nominal),
			Status:       string(trxData[i].Status),
		})
	}

	return
}
