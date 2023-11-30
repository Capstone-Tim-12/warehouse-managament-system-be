package model

type LoginRequest struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	UserId int    `json:"userId"`
	Name   string `json:"name"`
	Token  string `json:"token"`
}
