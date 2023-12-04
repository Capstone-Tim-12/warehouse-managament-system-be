package core

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (w *wrapper) CreateVA(ctx context.Context, req CreateVirtualAccountRequest) (resp CreateViartualAccountResponse, err error) {
	path := "/payment/va/create"

	fmt.Println(ctx, "[ CreateVA Request]", address+path, req)

	headers := getRequestHeaders(ctx)
	body, status, err := w.client.Post(ctx, path, headers, req)
	if err != nil {
		err = fmt.Errorf("[Core] CreateVA error: %v", err.Error())
		return
	}

	if status != http.StatusOK {
		err = fmt.Errorf("[Core] CreateVA return non 200 http status code. got %d", status)
		return
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		err = fmt.Errorf("[Core] Unmarshal Response Error %v", err.Error())
	}

	fmt.Println(ctx, "[CreateVA Response]", address+path, resp)

	return
}

func (w *wrapper) GetBank(ctx context.Context) (resp GetBankResponse, err error) {
	path := "/payment/va/bank"

	fmt.Println(ctx, "[GetBank Request]", address+path)

	headers := getRequestHeaders(ctx)
	body, status, err := w.client.Get(ctx, path, headers)
	if err != nil {
		err = fmt.Errorf("[Core] GetBank error: %v", err.Error())
		return
	}

	if status != http.StatusOK {
		err = fmt.Errorf("[Core] GetBank return non 200 http status code. got %d", status)
		return
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		err = fmt.Errorf("[Core] Unmarshal Response Error %v", err.Error())
	}

	fmt.Println(ctx, "[GetBank Response]", address+path, resp)

	return
}

func (w *wrapper) CheckPayment(ctx context.Context, paymentId string) (resp CheckPaymentResponse, err error) {
	path := "/payment/va/" + paymentId

	fmt.Println(ctx, "[CheckPayment Request]", address+path)

	headers := getRequestHeaders(ctx)
	body, status, err := w.client.Get(ctx, path, headers)
	if err != nil {
		err = fmt.Errorf("[Core] CheckPayment error: %v", err.Error())
		return
	}

	if status != http.StatusOK {
		err = fmt.Errorf("[Core] CheckPayment return non 200 http status code. got %d", status)
		return
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		err = fmt.Errorf("[Core] Unmarshal Response Error %v", err.Error())
	}

	fmt.Println(ctx, "[CheckPayment Response]", address+path, resp)

	return
}
