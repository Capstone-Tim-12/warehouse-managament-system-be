package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/regiondb"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/userdb"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/http/core"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/user/model"
	customError "github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/generate"
)

type defaultUser struct {
	regionRepo regiondb.RegionRepository
	userRepo   userdb.UserRepository
	coreRepo   core.CoreWrapper
}

func NewUserUsecase(regionRepo regiondb.RegionRepository, userRepo userdb.UserRepository, coreRepo core.CoreWrapper) *defaultUser {
	return &defaultUser{
		regionRepo: regionRepo,
		userRepo:   userRepo,
		coreRepo:   coreRepo,
	}
}

func (s *defaultUser) GetAllProvince(ctx context.Context) (resp []model.RegionResponse, err error) {
	data, err := s.regionRepo.FindAllProvince(ctx)
	if err != nil {
		err = errors.New("failed to get all province")
		return
	}
	for i := 0; i < len(data); i++ {
		resp = append(resp, model.RegionResponse{
			Id:   data[i].ID,
			Name: data[i].Name,
		})
	}
	return
}

func (s *defaultUser) GetRegencyByProvinceId(ctx context.Context, id string) (resp []model.RegionResponse, err error) {
	_, err = s.regionRepo.GetProvinceById(ctx, id)
	if err != nil {
		err = errors.New("province not found")
		return
	}

	data, err := s.regionRepo.FindRegencyByProvinceId(ctx, id)
	if err != nil {
		err = errors.New("failed to get data regency")
		return
	}

	for i := 0; i < len(data); i++ {
		resp = append(resp, model.RegionResponse{
			Id:   data[i].ID,
			Name: data[i].Name,
		})
	}
	return
}

func (s *defaultUser) GetDistricByRegencyId(ctx context.Context, id string) (resp []model.RegionResponse, err error) {
	_, err = s.regionRepo.GetRegencyById(ctx, id)
	if err != nil {
		err = errors.New("regency not found")
		return
	}

	data, err := s.regionRepo.FindDistrictByRegencyId(ctx, id)
	if err != nil {
		err = errors.New("failed to get data distric")
		return
	}

	for i := 0; i < len(data); i++ {
		resp = append(resp, model.RegionResponse{
			Id:   data[i].ID,
			Name: data[i].Name,
		})
	}
	return
}

func (s *defaultUser) RegisterData(ctx context.Context, req model.RegisterDataRequest) (err error) {
	userData, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		fmt.Println("Error getting Email", err.Error())
		err = customError.ErrNotFound
		return
	}

	_, err = s.regionRepo.GetProvinceById(ctx, req.ProvinceID)
	if err != nil {
		fmt.Println("Error getting province id", err.Error())
		err = customError.ErrNotFound
		return
	}
	_, err = s.regionRepo.GetRegencyById(ctx, req.RegencyID)
	if err != nil {
		fmt.Println("Error getting regency id", err.Error())
		err = customError.ErrNotFound
		return
	}
	_, err = s.regionRepo.GetDistrictById(ctx, req.DistrictID)
	if err != nil {
		fmt.Println("Error getting regency id", err.Error())
		err = customError.ErrNotFound
		return
	}

	createUserData := userdb.UserDetail{
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
		err = errors.New("internal error create user data")
		fmt.Println("Internal error create user data")
		return
	}

	userData.IsVerifyIdentity = true
	err = s.userRepo.UpdateUser(ctx, tx, userData)
	if err != nil {
		tx.Rollback()
		fmt.Println("error update user data")
		err = errors.New("failed to update data")
		return
	}

	tx.Commit()
	return
}

func (s *defaultUser) UserRegister(ctx context.Context, req model.RegisterUserRequest) (resp model.RegisterUserResponse, err error) {
	userdata, _ := s.userRepo.GetUserByEmail(ctx, req.Email)
	if userdata.Email != "" {
		err = errors.New("email already exists")
		fmt.Println("email already exists")
		return
	}
	passwordByrpt := HashPassword(req.Password)
	createUser := userdb.User{
		Username: req.Username,
		Password: passwordByrpt,
		Email:    req.Email,
	}
	tx := s.userRepo.BeginTrans(ctx)
	err = s.userRepo.CreateUser(ctx, tx, &createUser)
	if err != nil {
		tx.Rollback()
		err = errors.New("failed create data user")
		fmt.Println("failed create data user")
		return
	}
	tx.Commit()

	otpReq := model.OtpRequest{
		Email: req.Email,
	}

	err = s.ResendOtp(ctx, otpReq)
	if err != nil {
		s.userRepo.DeleteUser(ctx, &createUser)
		return
	}

	resp.Email = req.Email
	return
}

func (s *defaultUser) ResendOtp(ctx context.Context, req model.OtpRequest) (err error) {
	userData, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		fmt.Println("Error getting Email", err.Error())
		err = customError.ErrNotFound
		return
	}

	if userData.IsVerifyAccount {
		fmt.Println("your account has been verified")
		err = errors.New("your account has been verified")
		return err
	}

	otpMessage := generate.GenerateOTP()
	if err != nil {
		fmt.Println("failed to generate otp: ", err.Error())
		err = errors.New("failed to generate otp")
		return err
	}

	utilityData := core.SetUtilityRequest{
		Key:      userData.Email,
		Value:    otpMessage,
		Duration: 180,
	}

	_, err = s.coreRepo.SetUtility(ctx, utilityData)
	if err != nil {
		fmt.Println("failed to set utility: ", err.Error())
		err = errors.New("failed to set utility")
		return err
	}

	emailBody, _ := generate.GenerateEmailOTP(userData.Username, otpMessage)
	emailResponse := core.SendEmailRequest{
		To:       userData.Email,
		FromName: "DigiHouse Indonesia",
		Title:    "Verifikasi OTP untuk Akun Anda",
		Message:  emailBody,
	}

	_, err = s.coreRepo.SendEmail(ctx, emailResponse)
	if err != nil {
		fmt.Println("failed to send email: ", err.Error())
		err = errors.New("failed to send email")
		return err
	}

	return
}

func (s *defaultUser) Login(ctx context.Context, req model.LoginRequest) (resp model.LoginResponse, err error) {
	user, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		err = errors.New("email not found")
		return
	}

	err = ComparePassword(user.Password, req.Password)
	if err != nil {
		err = errors.New("invalid password")
		return
	}

	token := CreateToken(int(user.ID), string(user.Role))

	resp = model.LoginResponse{
		UserId: user.ID,
		Name:   user.Username,
		Token:  token,
	}
	return
}

func (s *defaultUser) VerificationUser(ctx context.Context, req model.VerificationUserRequest) (err error) {
	respData, err := s.coreRepo.GetUtilityData(ctx, req.Email)
	if err != nil {
		err = errors.New("failed verification otp")
		fmt.Println("timeout request", err.Error())
		return
	}

	if respData.Code != http.StatusOK {
		err = customError.ErrOTPWrong
		fmt.Println("otp is invalid")
		return
	}
	if req.Otp != respData.Data.Value {
		err = customError.ErrOTPWrong
		fmt.Println("otp is wrong")
		return
	}

	userData, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		fmt.Println("error getting user", err.Error())
		err = customError.ErrNotFound
		return
	}

	tx := s.userRepo.BeginTrans(ctx)
	userData.IsVerifyAccount = true
	err = s.userRepo.UpdateUser(ctx, tx, userData)
	if err != nil {
		tx.Rollback()
		fmt.Println("error updating user", err.Error())
		err = errors.New("failed to verification otp")
		return
	}

	tx.Commit()
	return
}
