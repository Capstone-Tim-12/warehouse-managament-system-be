package user

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/entity"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/user/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
)

func (s *defaultUser) UserRegister(ctx context.Context, req model.RegisterUserRequest, long, lat float64) (resp model.RegisterUserResponse, err error) {
	userdata, _ := s.userRepo.GetUserByEmail(ctx, req.Email)
	if userdata.Email != "" {
		err = errors.New(http.StatusConflict, "email already exists")
		fmt.Println("email already exists")
		return
	}

	separator := strings.LastIndex(req.Email, "@")
	host := req.Email[separator+1:]
	mxRecords, err := net.LookupMX(host)
	if err != nil {
		fmt.Println("error Lookup MX: ", err.Error())
		err = errors.New(http.StatusBadRequest, "bad email address")
		return
	}

	if len(mxRecords) == 0 {
		fmt.Println("no MX records found")
		err = errors.New(http.StatusBadRequest, "email address not valid")
		return
	}

	passwordByrpt := HashPassword(req.Password)
	createUser := entity.User{
		Username:  req.Username,
		Password:  passwordByrpt,
		Email:     req.Email,
		Role:      entity.RoleUser,
		Longitude: long,
		Latitude:  lat,
	}
	tx := s.userRepo.BeginTrans(ctx)
	err = s.userRepo.CreateUser(ctx, tx, &createUser)
	if err != nil {
		tx.Rollback()
		err = errors.New(http.StatusInternalServerError, "failed create data user")
		fmt.Println("failed create data user")
		return
	}

	err = s.sendEmailOtp(ctx, createUser)
	if err != nil {
		fmt.Println("failed send email otp")
		tx.Rollback()
		return
	}

	resp.Email = req.Email
	tx.Commit()
	return
}

func (s *defaultUser) RegisterData(ctx context.Context, req model.RegisterDataRequest) (err error) {
	userData, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		fmt.Println("Error getting Email", err.Error())
		err = errors.New(http.StatusNotFound, "email not found")
		return
	}

	if userData.IsVerifyIdentity {
		fmt.Println("user has verify identity")
		err = errors.New(http.StatusBadRequest, "user has verify identity")
		return
	}

	_, err = s.regionRepo.GetDistrictById(ctx, req.DistrictID)
	if err != nil {
		fmt.Println("Error getting regency id", err.Error())
		err = errors.New(http.StatusNotFound, "district not found")
		return
	}

	createUserData := entity.UserDetail{
		NIK:          req.NIK,
		Address:      req.Address,
		Country:      "Indonesia",
		FullName:     req.FullName,
		Gender:       req.Gender,
		PlaceOfBirth: req.PlaceBirth,
		DateBirth:    req.DateBirth,
		Work:         req.Work,
		Citizenship:  req.Citizenship,
		UserID:       userData.ID,
		DistrictID:   req.DistrictID,
	}

	tx := s.userRepo.BeginTrans(ctx)
	err = s.userRepo.CreateDetail(ctx, tx, &createUserData)
	if err != nil {
		tx.Rollback()
		err = errors.New(http.StatusInternalServerError, "error create user data")
		fmt.Println("Internal error create user data")
		return
	}

	userData.IsVerifyIdentity = true
	err = s.userRepo.UpdateUser(ctx, tx, userData)
	if err != nil {
		tx.Rollback()
		fmt.Println("error update user data")
		err = errors.New(http.StatusInternalServerError, "failed to update data")
		return
	}

	tx.Commit()
	return
}
