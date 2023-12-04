package payment

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/entity"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/http/core"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/payment/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/constrans"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/generate"
)

func (s *defaultPayment) PaymentCheckout(ctx context.Context, req model.PaymentRequest) (resp model.PaymentResponse, err error) {
	_, err = s.paymentRepo.GetPaymentMethodById(ctx, req.PaymentMethodId)
	if err != nil {
		fmt.Println("error getting payment method: ", err.Error())
		err = errors.New(http.StatusNotFound, "payment method not found")
		return
	}

	instalmentData, err := s.paymentRepo.GetInstalmentById(ctx, req.InstalmentId)
	if err != nil {
		fmt.Println("error getting instalment: ", err.Error())
		err = errors.New(http.StatusNotFound, "instalment not found")
		return
	}

	if instalmentData.OngoingInstalment.ID != 0 {
		fmt.Println("Ongoing instalment is not empty")
		err = errors.New(http.StatusBadRequest, "If you have made a payment request, please contact admin")
		return
	}

	if instalmentData.Status != entity.Unpaid {
		fmt.Println("payment for this transaction is not permitted")
		err = errors.New(http.StatusBadRequest, "Payment for this transaction is not permitted")
		return
	}

	if instalmentData.DueDate.Before(time.Now()) {
		fmt.Println("payment is due date")
		err = errors.New(http.StatusBadRequest, "payment is due please contant admin")
		return
	}

	switch req.PaymentMethodId {
	case constrans.PaymentVirtualAccount:
		var vaData model.VaDataRequest
		json.Unmarshal([]byte(req.Data), &vaData)
		if vaData.BankCode == "" {
			fmt.Println("bank code is empty")
			err = errors.New(http.StatusBadRequest, "bank code is empty")
			return
		}

		reqVa := core.CreateVirtualAccountRequest{
			ExternalID:     generate.GenerateExternalId(15),
			BankCode:       strings.ToUpper(vaData.BankCode),
			Name:           "Digihouse Payment",
			IsSingleUse:    true,
			IsClosed:       true,
			ExpectedAmount: int(instalmentData.Nominal),
			ExpirationDate: time.Now().Add(time.Duration(24) * time.Hour),
		}
		vaResp, errRes := s.coreWrapper.CreateVA(ctx, reqVa)
		if errRes != nil {
			fmt.Println("error creating VA: ", errRes.Error())
			err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}

		if vaResp.Code != http.StatusOK {
			fmt.Println("error creating VA: ", vaResp.Code)
			err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}

		ongoingReq := entity.OngoingInstalment{
			InstalmentID:    req.InstalmentId,
			Instalment:      entity.Instalment{},
			PaymentMethodID: req.PaymentMethodId,
			PaymentMethod:   entity.PaymentMethod{},
			XPayment:        vaResp.Data.ExternalID,
			AccountNumber:   vaResp.Data.AccountNumber,
			TotalPayment:    float64(vaResp.Data.ExpectedAmount),
			BankCode:        vaResp.Data.BankCode,
			Expired:         vaResp.Data.ExpirationDate,
		}
		err = s.paymentRepo.CreateOngoingInstalment(ctx, &ongoingReq)
		if err != nil {
			fmt.Println("error create ongoing instalment: ", err.Error())
			err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}

		vaResponse := model.VaDataResponse{
			XpaymentId:           vaResp.Data.ExternalID,
			VirtualAccountName:   vaResp.Data.Name,
			VirtualAccountNumber: vaResp.Data.AccountNumber,
			BankCode:             vaResp.Data.BankCode,
			Nominal:              vaResp.Data.ExpectedAmount,
			ExpiredAt:            vaResp.Data.ExpirationDate,
		}

		respData, _ := json.Marshal(vaResponse)
		resp.PaymentInfo = string(respData)

	case constrans.PaymentDebitCredit:
		err = errors.New(http.StatusForbidden, "credit card not available")
		return

	default:
		err = errors.New(http.StatusForbidden, "payment method is not recognized")
		return
	}

	return
}

func (s *defaultPayment) VaCallback(ctx context.Context, req model.VaCallbackRequest) (err error) {
	ongoingData, err := s.paymentRepo.FindOngoingInstalmentByXpayment(ctx, req.ExternalID)
	if err != nil {
		fmt.Println("error getting ongoing data: ", err.Error())
		err = errors.New(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}
	xenditResp, err := s.coreWrapper.CheckPayment(ctx, req.ExternalID)
	if err != nil {
		fmt.Println("error check payment", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	var instalmentStatus entity.InstalmentStatus
	var paymentTime time.Time
	if xenditResp.Data.CallbackVirtualAccountID != "" {
		instalmentStatus = entity.Paid
		paymentTime = xenditResp.Data.TransactionTimestamp
	} else {
		instalmentStatus = entity.Failed
	}

	tx := s.paymentRepo.BeginTrans(ctx)
	if instalmentStatus == entity.Paid {
		ongoingData.PaymentTime = &paymentTime
		err = s.paymentRepo.UpdateOngoingInstalment(ctx, tx, ongoingData)
		if err != nil {
			tx.Rollback()
			fmt.Println("error update payment", err.Error())
			err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
	}
	instalmentData, err := s.paymentRepo.GetInstalmentById(ctx, ongoingData.InstalmentID)
	if err != nil {
		tx.Rollback()
		fmt.Println("error get instalment", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusNotFound))
		return
	}
	instalmentData.Status = instalmentStatus
	err = s.paymentRepo.UpdateInstalment(ctx, tx, instalmentData)
	if err != nil {
		tx.Rollback()
		fmt.Println("error update instalment", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusNotFound))
		return
	}

	tx.Commit()
	return
}

func (s *defaultPayment) GetBankVa(ctx context.Context) (resp []model.VaBankResponse, err error) {
	data, err := s.coreWrapper.GetBank(ctx)
	if err != nil {
		fmt.Println("error getting bank", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	if data.Code != http.StatusOK {
		fmt.Println("failed to get bank")
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	for i := 0; i < len(data.Data); i++ {
		resp = append(resp, model.VaBankResponse{
			Name: data.Data[i].Name,
			Code: data.Data[i].Code,
		})
	}
	return
}
