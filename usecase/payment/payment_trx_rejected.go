package payment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/entity"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
)

func (s *defaultPayment) TransactionRejected(ctx context.Context, transactionId string) (err error) {
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
	trxData.Status = entity.Rejected
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
