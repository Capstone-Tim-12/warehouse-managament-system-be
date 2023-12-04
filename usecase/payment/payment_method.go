package payment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/payment/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
)

func (s *defaultPayment) GetListPaymentMethod(ctx context.Context) (resp []model.PaymentMethodResponse, err error) {
	data, err := s.paymentRepo.GetListPaymentMethod(ctx)
	if err != nil {
		fmt.Println("error getting payment method: ", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	for i := 0; i < len(data); i++ {
		resp = append(resp, model.PaymentMethodResponse{
			Id:    data[i].ID,
			Name:  data[i].Name,
			Image: data[i].Image,
		})
	}

	return
}
