package payment

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"strings"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/entity"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/payment/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/constrans"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/paginate"
)

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

func (s *defaultPayment) GetTransactionListDetail(ctx context.Context, transactionId string) (resp model.TrxListDetail, err error) {
	trxData, err := s.paymentRepo.GetTransactionDetailById(ctx, transactionId)
	if err != nil {
		fmt.Println("error getting transaction")
		err = errors.New(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}
	var image string
	if len(trxData.Warehouse.WarehouseImg) != 0 {
		image = trxData.Warehouse.WarehouseImg[0].Image
	}
	resp = model.TrxListDetail{
		WarehouseId:       trxData.WarehouseID,
		WarehouseName:     trxData.Warehouse.Name,
		WarehousePrice:    trxData.Warehouse.Price,
		WarehouseAdreess:  trxData.Warehouse.Address,
		WarehouseDistrict: trxData.Warehouse.District.Name,
		WarehouseRegency:  trxData.Warehouse.District.Regency.Name,
		WarehouseProvince: trxData.Warehouse.District.Regency.Province.Name,
		WarehouseImage:    image,
		Username:          trxData.User.Username,
		IsVerifyIdentity:  trxData.User.IsVerifyIdentity,
		RentalDuration:    trxData.Duration,
		PaymentScheme:     trxData.PaymentScheme.Scheme,
	}
	return
}

func (s *defaultPayment) GetListTrxUserDasboar(ctx context.Context, userId int, param paginate.Pagination) (resp []model.ListTrxUserDasboarResponse, count int64, err error) {
	trxData, count, err := s.paymentRepo.GetListTransactionByUserIdAndStatus(ctx, userId, entity.Approved, param)
	if err != nil {
		fmt.Println("error getting transaction")
		err = errors.New(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	var price float64 
	for i := 0; i < len(trxData); i++ {
		if strings.EqualFold(trxData[i].PaymentScheme.Scheme, constrans.PaymentSchemeAnnualy) {
			price = trxData[i].Warehouse.Price
		} else if strings.EqualFold(trxData[i].PaymentScheme.Scheme, constrans.PaymentSchemeMonthly) {
			price = trxData[i].Warehouse.Price / 12
		} else if strings.EqualFold(trxData[i].PaymentScheme.Scheme, constrans.PaymentSchemeWeekly) {
			price = trxData[i].Warehouse.Price / 52
		} else {
			fmt.Println("data payment scheme not supported")
			err = errors.New(http.StatusForbidden, "data payment scheme not supported")
			return
		}

		resp = append(resp, model.ListTrxUserDasboarResponse{
			TransactionId:  trxData[i].ID,
			RentalDuration: trxData[i].Duration,
			PaymentScheme:  trxData[i].PaymentScheme.Scheme,
			PaymentTotal:   math.Ceil(price) * float64(trxData[i].Duration),
		})
	}

	return
}
