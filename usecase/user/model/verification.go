package model

type VerificationUserRequest struct {
	Email string `json:"email" validate:"email,required"`
	Otp   string `json:"otp" validate:"required"`
}

type VerificationUserResponse struct {
	Email   string `json:"email"`
	VerfyId string `json:"verfyId"`
}
