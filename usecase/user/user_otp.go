package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/entity"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/http/core"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/user/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/constrans"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/generate"
)

func (s *defaultUser) ResendOtp(ctx context.Context, req model.OtpRequest) (err error) {
	userData, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		fmt.Println("Error getting Email", err.Error())
		err = errors.New(http.StatusNotFound, "email not found")
		return
	}

	err = s.sendEmailOtp(ctx, *userData)
	if err != nil {
		return
	}

	return
}

func (s *defaultUser) VerificationUser(ctx context.Context, req model.VerificationUserRequest) (resp model.VerificationUserResponse, err error) {
	respData, err := s.coreRepo.GetUtilityData(ctx, req.Email)
	if err != nil {
		err = errors.New(http.StatusInternalServerError, "failed verification otp")
		fmt.Println("timeout request to utility", err.Error())
		return
	}

	if respData.Code != http.StatusOK {
		err = errors.New(http.StatusBadRequest, "wrong verification code")
		fmt.Println("otp is invalid")
		return
	}
	if req.Otp != respData.Data.Value {
		err = errors.New(http.StatusBadRequest, "wrong verification code")
		fmt.Println("otp is wrong")
		return
	}

	userData, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		fmt.Println("error getting user", err.Error())
		err = errors.New(http.StatusNotFound, "email not found")
		return
	}

	tx := s.userRepo.BeginTrans(ctx)
	userData.IsVerifyAccount = true

	err = s.userRepo.UpdateUser(ctx, tx, userData)
	if err != nil {
		tx.Rollback()
		fmt.Println("error updating user", err.Error())
		err = errors.New(http.StatusInternalServerError, "failed to verification otp")
		return
	}

	verifyId := generate.GenerateRandomString(25)
	reqSet := core.SetUtilityRequest{
		Key:      constrans.KeyVerify + req.Email,
		Value:    verifyId,
		Duration: 180,
	}

	_, err = s.coreRepo.SetUtility(ctx, reqSet)
	if err != nil {
		fmt.Println("failed set utility", err.Error())
		err = errors.New(http.StatusInternalServerError, "failed to verification otp")
		return
	}

	resp.VerfyId = verifyId
	resp.Email = userData.Email
	tx.Commit()
	return
}

func (s *defaultUser) sendEmailOtp(ctx context.Context, userData entity.User) (err error) {
	otpMessage := generate.GenerateOTP()
	if err != nil {
		fmt.Println("failed to generate otp: ", err.Error())
		err = errors.New(http.StatusInternalServerError, "failed to generate otp")
		return
	}

	utilityData := core.SetUtilityRequest{
		Key:      userData.Email,
		Value:    otpMessage,
		Duration: 180,
	}

	_, err = s.coreRepo.SetUtility(ctx, utilityData)
	if err != nil {
		fmt.Println("failed to set utility: ", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
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
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	return
}
