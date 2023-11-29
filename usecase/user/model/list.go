package model

type UserListResponse struct {
	UserId   int    `json:"userId"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Photo    string `json:"photo"`
}
