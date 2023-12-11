package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/user/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
)

func (s *defaultUser) UpdateAdminDasboarUser(ctx context.Context, userId int, req model.UserSettingRequest) (err error) {
	userData, err := s.userRepo.GetUserById(ctx, userId)
	if err != nil {
		fmt.Println("Error getting user by Id", err.Error())
		err = errors.New(http.StatusNotFound, "user not found")
		return
	}

	if userData.Email != req.Email {
		userdata, _ := s.userRepo.GetUserByEmail(ctx, req.Email)
		if userdata.Email != "" {
			err = errors.New(http.StatusConflict, "email already exists")
			fmt.Println("email already exists")
			return
		}
	}

	userData.Email = req.Email
	userData.Username = req.Username
	if req.Password != "" {
		pass := HashPassword(req.Password)
		userData.Password = pass
	}
	tx := s.userRepo.BeginTrans(ctx)
	err = s.userRepo.UpdateUser(ctx, tx, userData)
	if err != nil {
		tx.Rollback()
		fmt.Println("err updating user by Id", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	tx.Commit()
	return
}
