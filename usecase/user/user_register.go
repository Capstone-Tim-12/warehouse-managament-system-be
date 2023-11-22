package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/entity"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/user/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
)

func (s *defaultUser) UserRegister(ctx context.Context, req model.RegisterUserRequest) (resp model.RegisterUserResponse, err error) {
	userdata, _ := s.userRepo.GetUserByEmailUsername(ctx, req.Email, req.Username)
	if userdata.Email != "" {
		err = errors.New(http.StatusConflict, "email or username already exists")
		fmt.Println("email or username already exists")
		return
	}

	passwordByrpt := HashPassword(req.Password)
	createUser := entity.User{
		Username: req.Username,
		Password: passwordByrpt,
		Email:    req.Email,
		Role:     entity.RoleUser,
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

	_, err = s.regionRepo.GetProvinceById(ctx, req.ProvinceID)
	if err != nil {
		fmt.Println("Error getting province id", err.Error())
		err = errors.New(http.StatusNotFound, "province not found")
		return
	}
	_, err = s.regionRepo.GetRegencyById(ctx, req.RegencyID)
	if err != nil {
		fmt.Println("Error getting regency id", err.Error())
		err = errors.New(http.StatusNotFound, "regency not found")
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
		ProvinceID:   req.ProvinceID,
		RegencyID:    req.RegencyID,
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