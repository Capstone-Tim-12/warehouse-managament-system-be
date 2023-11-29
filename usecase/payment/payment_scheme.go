package payment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/payment/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
)

func (s *defaultPayment) GetPaymentScheme(ctx context.Context) (resp []model.SchemeResponse, err error) {
	schemeData, err := s.paymentRepo.FindPaymentScheme(ctx)
	if err != nil {
		fmt.Println("error finding scheme: ", err.Error())
		err = errors.New(http.StatusNotFound, "scheme not found")
		return
	}

	for _, scheme := range schemeData {
		resp = append(resp, model.SchemeResponse{
			ID:     scheme.ID,
			Scheme: scheme.Scheme,
		})
	}
	return
}
