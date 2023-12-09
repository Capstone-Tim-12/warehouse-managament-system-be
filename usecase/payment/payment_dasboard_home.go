package payment

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/entity"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/payment/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
)

func (s *defaultPayment) GetTotalPaymentDasboard(ctx context.Context) (resp model.GetTotalPayment, err error) {
	warehouseAvailable, err := s.warehouseRepo.GetTotalWarehouseByStatus(ctx, entity.Available)
	if err != nil {
		fmt.Println("failed to get totalWarehouse by status: ", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	warehouseNotAvailable, err := s.warehouseRepo.GetTotalWarehouseByStatus(ctx, entity.NotAvailable)
	if err != nil {
		fmt.Println("failed to get totalWarehouse by status: ", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	totalPayment, err := s.paymentRepo.GetTotalPayment(ctx)
	if err != nil {
		fmt.Println("failed to get payment: ", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	resp = model.GetTotalPayment{
		TotalWarehouseAvailabe:     warehouseAvailable,
		TotalWarehouseNotAvailable: warehouseNotAvailable,
		TotalPayment:               totalPayment,
	}

	return
}

func (s *defaultPayment) GetStatistictPaymentOnYear(ctx context.Context) (resp []model.StatiscticPayment, err error) {
	for i := -4; i <= 0; i++ {
		currentTime := time.Now()
		yearsAgo := currentTime.Add(time.Duration(i) * 365 * 24 * time.Hour)
		totalPayment, errRes := s.paymentRepo.GetTotalPaymentOnYear(ctx, yearsAgo.Year())
		if errRes != nil {
			fmt.Println("error getting total payment on year: ", errRes.Error())
			err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}

		resp = append(resp, model.StatiscticPayment{
			Year:         yearsAgo.Year(),
			TotalPayment: totalPayment,
		})
	}

	return
}
