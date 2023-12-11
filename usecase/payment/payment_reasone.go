package payment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/payment/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
)

func (s *defaultPayment) GetReasoneList(ctx context.Context) (resp []model.ReasoneData, err error) {
	data, err := s.paymentRepo.GetReason(ctx)
	if err != nil {
		fmt.Println("error getting Reasone")
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	for i := 0; i < len(data); i++ {
		resp = append(resp, model.ReasoneData{
			Id:   data[i].ID,
			Name: data[i].Name,
		})
	}

	return
}