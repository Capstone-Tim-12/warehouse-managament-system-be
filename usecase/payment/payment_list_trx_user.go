package payment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/payment/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/paginate"
)

func (s *defaultPayment) GetListTransactionByUserId(ctx context.Context, userId int) (resp []model.ListTrxUser, err error) {
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

func (s *defaultPayment) GetAllTransaction(ctx context.Context, param paginate.PaginationTrx) (resp []model.ListAllTrxResponse, count int64, err error) {
	trxData, count, err := s.paymentRepo.GetListTransactionData(ctx, param)
	if err != nil {
		fmt.Println("error getting transaction: ", err.Error())
		err = errors.New(http.StatusInternalServerError, "error getting transaction")
		return
	}

	for i := 0; i < len(trxData); i++ {
		resp = append(resp, model.ListAllTrxResponse{
			TransactionId: trxData[i].ID,
			UserId:        trxData[i].UserID,
			Username:      trxData[i].User.Username,
			RegencyId:     trxData[i].Warehouse.District.Regency.ID,
			RegencyName:   trxData[i].Warehouse.District.Regency.Name,
			ProvinceId:    trxData[i].Warehouse.District.Regency.Province.ID,
			ProvinceName:  trxData[i].Warehouse.District.Regency.Province.Name,
			WarehouseName: trxData[i].Warehouse.Name,
			WarehouseId:   trxData[i].Warehouse.ID,
			Duration:      trxData[i].Duration,
			PaymentScheme: trxData[i].PaymentScheme.Scheme,
			Status:        string(trxData[i].Status),
		})
	}
	return
}