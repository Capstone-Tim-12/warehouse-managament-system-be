package model

type OtpRequest struct {
	Email string `json:"email" validate:"required,email"`
}
