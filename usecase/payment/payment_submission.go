package payment

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/entity"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/payment/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/constrans"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
	"github.com/spf13/cast"
)

func (s *defaultPayment) SubmissionWarehouse(ctx context.Context, userId int, req model.SubmissionRequest) (err error) {
	warehouseData, err := s.warehouseRepo.FindWarehouseByIdOnly(ctx, cast.ToString(req.WarehouseId))
	if err != nil {
		fmt.Println("error finding warehouse: ", err.Error())
		err = errors.New(http.StatusNotFound, "warehouse not found")
		return
	}

	if warehouseData.Status == entity.NotAvailable {
		fmt.Println("warehouse not available")
		err = errors.New(http.StatusBadRequest, "warehouse not available")
		return
	}

	schemeData, err := s.paymentRepo.FindPaymentSchemeById(ctx, req.PaymentSchemeId)
	if err != nil {
		fmt.Println("error finding payment scheme: ", err.Error())
		err = errors.New(http.StatusNotFound, "payment scheme not found")
		return
	}

	if req.DateEntry.Before(time.Now()) {
		fmt.Println("input date cannot be less than the current date")
		err = errors.New(http.StatusBadRequest, "date cannot be less than the current date")
		return
	}

	var dateOut time.Time
	if strings.EqualFold(schemeData.Scheme, constrans.PaymentSchemeAnnualy) {
		dateOut = time.Now().AddDate(req.Duration, 0, 0)
	} else if strings.EqualFold(schemeData.Scheme, constrans.PaymentSchemeMonthly) {
		dateOut = time.Now().AddDate(0, req.Duration, 0)
	} else if strings.EqualFold(schemeData.Scheme, constrans.PaymentSchemeWeekly) {
		dateOut = time.Now().AddDate(0, 0, req.Duration*7)
	} else {
		fmt.Println("data payment scheme not supported")
		err = errors.New(http.StatusForbidden, "data payment scheme not supported")
		return
	}

	reqTransaction := entity.Transaction{
		DateEntry:       req.DateEntry,
		DateOut:         dateOut,
		UserID:          userId,
		PaymentSchemeID: req.PaymentSchemeId,
		Duration:        req.Duration,
		Status:          entity.Submission,
		WarehouseID:     req.WarehouseId,
	}
	tx := s.paymentRepo.BeginTrans(ctx)
	err = s.paymentRepo.CreateTransaction(ctx, tx, &reqTransaction)
	if err != nil {
		tx.Rollback()
		fmt.Println("create transaction failed: ", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	warehouseData.Status = entity.NotAvailable
	err = s.warehouseRepo.UpdateWarehouse(ctx, tx, warehouseData)
	if err != nil {
		tx.Rollback()
		fmt.Println("failed update warehouse: ", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	tx.Commit()
	return
}