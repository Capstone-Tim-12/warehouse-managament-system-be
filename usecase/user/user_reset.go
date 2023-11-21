package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/user/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/constrans"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
)

func (s *defaultUser) ResetPassword(ctx context.Context, req model.ResetPasswordRequest) (err error) {
	key := constrans.KeyVerify + req.Email
	respData, err := s.coreRepo.GetUtilityData(ctx, key)
	if err != nil {
		err = errors.New(http.StatusInternalServerError, "failed verification otp")
		fmt.Println("timeout request", err.Error())
		return
	}

	if respData.Code != http.StatusOK {
		err = errors.New(http.StatusBadRequest, "invalid verify id")
		fmt.Println("verifyId is invalid")
		return
	}
	if req.VerifyId != respData.Data.Value {
		err = errors.New(http.StatusBadRequest, "verify id is wrong")
		fmt.Println("verifyId is wrong")
		return
	}
	userData, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		err = errors.New(http.StatusNotFound, "email not found")
		return
	}

	passHas := HashPassword(req.NewPassword)
	userData.Password = passHas
	tx := s.userRepo.BeginTrans(ctx)
	err = s.userRepo.UpdateUser(ctx, tx, userData)
	if err != nil {
		tx.Rollback()
		fmt.Println("error update user")
		err = errors.New(http.StatusInternalServerError, "error reset password")
		return
	}

	tx.Commit()
	return
}
