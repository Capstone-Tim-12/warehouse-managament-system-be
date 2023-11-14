package model

type VerificationUserRequest struct {
	Email string `json:"email"`
	Otp   string `json:"otp"`
}

type VerificationUserResponse struct {
	Email   string `json:"email"`
	VerfyId string `json:"verfyId"`
}
