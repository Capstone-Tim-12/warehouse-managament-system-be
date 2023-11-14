package model

type ResetPasswordRequest struct {
	VerifyId string `json:"verifyId"`
	Email    string `json:"email"`
	NewPassword string `json:"newPassword"`
}
