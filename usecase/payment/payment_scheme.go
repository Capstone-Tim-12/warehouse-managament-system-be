package payment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/payment/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
)

func (s *defaultPayment) GetPaymentScheme(ctx context.Context, id int) (resp []model.SchemeResponse, err error) {
	_, err = s.userRepo.GetUserById(ctx, id)
	if err != nil {
		fmt.Println("error finding user: ", err.Error())
		err = errors.New(http.StatusNotFound, "user not found")
		return
	}

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
