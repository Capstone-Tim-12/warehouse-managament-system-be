package payment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/payment/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
)

func (s *defaultPayment) GetTransactionDetailDasboardUser(ctx context.Context, transactionId string) (resp model.TransactionDetailUser, err error) {
	trxData, err := s.paymentRepo.GetTransactionUserDetailByTransactionId(ctx, transactionId)
	if err != nil {
		fmt.Println("error getting transaction: ", err.Error())
		err = errors.New(http.StatusNotFound, "transaction not found")
		return
	}

	resp = model.TransactionDetailUser{
		WarehouseName:    trxData.Warehouse.Name,
		Username:         trxData.User.Username,
		IsVerifyIdentity: trxData.User.IsVerifyIdentity,
		Address:          trxData.User.UserDetail.Address,
		DistricName:      trxData.User.UserDetail.District.Name,
		RegencyName:      trxData.User.UserDetail.District.Regency.Name,
		Duration:         trxData.Duration,
		PaymentScheme:    trxData.PaymentScheme.Scheme,
		EntryDate:        trxData.DateEntry,
		OutDate:          trxData.DateOut,
	}

	for i := 0; i < len(trxData.Instalment); i++ {
		data := model.InstalmentList{
			InstalmentId: trxData.Instalment[i].ID,
			DueDate:      trxData.Instalment[i].DueDate,
			Nominal:      trxData.Instalment[i].Nominal,
			Status:       string(trxData.Instalment[i].Status),
		}

		if trxData.Instalment[i].OngoingInstalment != nil {
			if trxData.Instalment[i].OngoingInstalment.PaymentTime != nil {
				data.PaymentTime = trxData.Instalment[i].OngoingInstalment.PaymentTime
			}
			data.PaymentName = fmt.Sprintf("%v %v", trxData.Instalment[i].OngoingInstalment.PaymentMethod.Name, trxData.Instalment[i].OngoingInstalment.BankCode)
		}

		resp.TotalPayment = resp.TotalPayment + trxData.Instalment[i].Nominal
		resp.Instalment = append(resp.Instalment, data)
	}

	return
}
