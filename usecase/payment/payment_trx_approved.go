package payment

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"strings"
	"time"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/entity"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/constrans"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
)

func (s *defaultPayment) TransactionApproved(ctx context.Context, transactionId string) (err error) {
	trxData, err := s.paymentRepo.GetTransactionById(ctx, transactionId)
	if err != nil {
		fmt.Println("Error getting transaction: ", err.Error())
		err = errors.New(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	if trxData.Status != entity.Submission {
		fmt.Println("status transaction not submission")
		err = errors.New(http.StatusBadRequest, "status transaction not submission")
		return
	}

	tx := s.paymentRepo.BeginTrans(ctx)
	var dueDate time.Time
	var nominal float64
	for i := 0; i < trxData.Duration; i++ {
		if strings.EqualFold(trxData.PaymentScheme.Scheme, constrans.PaymentSchemeAnnualy) {
			dueDate = trxData.DateEntry.AddDate(i, 0, 0)
			nominal = trxData.Warehouse.Price
		} else if strings.EqualFold(trxData.PaymentScheme.Scheme, constrans.PaymentSchemeMonthly) {
			dueDate = time.Now().AddDate(0, i, 0)
			nominal = math.Ceil(trxData.Warehouse.Price / 12)
		} else if strings.EqualFold(trxData.PaymentScheme.Scheme, constrans.PaymentSchemeWeekly) {
			dueDate = time.Now().AddDate(0, 0, i*7)
			nominal = math.Ceil(trxData.Warehouse.Price / 52)
		} else {
			fmt.Println("data payment scheme not supported")
			err = errors.New(http.StatusForbidden, "data payment scheme not supported")
			return
		}
		reqInstalment := entity.Instalment{
			TransactionID: trxData.ID,
			Nominal:       nominal,
			DueDate:       dueDate,
			Status:        entity.Unpaid,
		}
		err = s.paymentRepo.CreateInstalment(ctx, tx, &reqInstalment)
		if err != nil {
			tx.Rollback()
			fmt.Println("Error creating transaction instalment: ", err.Error())
			err = errors.New(http.StatusInternalServerError, "error creating transaction")
			return
		}
	}

	trxData.Status = entity.Approved
	err = s.paymentRepo.UpdateTransaction(ctx, tx, trxData)
	if err != nil {
		tx.Rollback()
		fmt.Println("Error updating transaction")
		err = errors.New(http.StatusInternalServerError, "error updating transaction")
		return
	}

	tx.Commit()
	return
}
