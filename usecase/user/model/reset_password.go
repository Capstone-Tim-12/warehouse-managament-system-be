package model

type ResetPasswordRequest struct {
	VerifyId    string `json:"verifyId" validate:"required"`
	Email       string `json:"email" validate:"email,required"`
	NewPassword string `json:"newPassword" validate:"min=6"`
}
