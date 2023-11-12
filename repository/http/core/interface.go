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
}