package user

import (
	"context"
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/user/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
)

func (s *defaultUser) Login(ctx context.Context, req model.LoginRequest) (resp model.LoginResponse, err error) {
	user, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		err = errors.New(http.StatusNotFound, "email not found")
		return
	}

	err = ComparePassword(user.Password, req.Password)
	if err != nil {
		err = errors.New(http.StatusBadRequest, "invalid password")
		return
	}

	role := string(user.Role) 
	token := CreateToken(int(user.ID), role)

	resp = model.LoginResponse{
		UserId: user.ID,
		Name:   user.Username,
		Token:  token,
	}
	return
}
