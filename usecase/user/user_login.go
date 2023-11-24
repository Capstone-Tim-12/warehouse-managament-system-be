package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/user/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
)

func (s *defaultUser) Login(ctx context.Context, req model.LoginRequest, lat, long float64) (resp model.LoginResponse, err error) {
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

	user.Latitude = lat
	user.Longitude = long
	
	tx := s.userRepo.BeginTrans(ctx)
	err = s.userRepo.UpdateUser(ctx, tx, user)
	if err != nil {
		tx.Rollback()
		fmt.Println("error updating user login: ", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	role := string(user.Role) 
	token := CreateToken(int(user.ID), role)

	tx.Commit()
	resp = model.LoginResponse{
		UserId: user.ID,
		Name:   user.Username,
		Token:  token,
	}
	return
}
