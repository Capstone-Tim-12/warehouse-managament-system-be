package core

import (
	"context"
	"mime/multipart"
)

type CoreWrapper interface {
	GetUtilityData(ctx context.Context, key string) (resp GetUtilityResponse, err error)
	SetUtility(ctx context.Context, req SetUtilityRequest) (resp SetUtilityResponse, err error)
	SendEmail(ctx context.Context, req SendEmailRequest) (resp SendEmailResponse, err error)
	UploadImage(ctx context.Context, req *multipart.FileHeader) (resp UploadImageResponse, err error)
	CreateVA(ctx context.Context, req CreateVirtualAccountRequest) (resp CreateViartualAccountResponse, err error)
	GetBank(ctx context.Context) (resp GetBankResponse, err error)
	CheckPayment(ctx context.Context, paymentId string) (resp CheckPaymentResponse, err error)
}