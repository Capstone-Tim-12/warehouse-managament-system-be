package payment

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"strings"
	"time"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/payment/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/constrans"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/paginate"
)

func (s *defaultPayment) GetListTranscationByWarehouseId(ctx context.Context, warehouseId int, param paginate.Pagination) (resp []model.ListTransactionWarehouseDasboard, count int64, err error) {
	wrData, count, err := s.paymentRepo.GetTransactionDetailByWarehouseId(ctx, warehouseId, param)
	if err != nil {
		fmt.Println("error getting transaction detail: ", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	var price float64
	for i := 0; i < len(wrData); i++ {
		if strings.EqualFold(wrData[i].PaymentScheme.Scheme, constrans.PaymentSchemeAnnualy) {
			price = wrData[i].Warehouse.Price
		} else if strings.EqualFold(wrData[i].PaymentScheme.Scheme, constrans.PaymentSchemeMonthly) {
			price = wrData[i].Warehouse.Price / 12
		} else if strings.EqualFold(wrData[i].PaymentScheme.Scheme, constrans.PaymentSchemeWeekly) {
			price = wrData[i].Warehouse.Price / 52
		} else {
			fmt.Println("data payment scheme not supported")
			err = errors.New(http.StatusForbidden, "data payment scheme not supported")
			return
		}

		data := model.ListTransactionWarehouseDasboard{
			TransactionId:   wrData[i].ID,
			Username:        wrData[i].User.Username,
			UserRegencyName: wrData[i].User.UserDetail.District.Regency.Name,
			Nominal:         math.Ceil(price * float64(wrData[i].Duration)),
			Status:          model.PaymentActive,
		}

		if wrData[i].DateOut.Before(time.Now()) {
			data.Status = model.PaymentFinish
		}
		resp = append(resp, data)
	}
	return
}
