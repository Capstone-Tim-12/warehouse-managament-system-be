package model

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserId int    `json:"userId"`
	Name   string `json:"name"`
	Token  string `json:"token"`
}
