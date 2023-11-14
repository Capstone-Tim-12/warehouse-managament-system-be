package model

type VerificationUserRequest struct {
	Email string `json:"email"`
	Otp   string `json:"otp"`
}
