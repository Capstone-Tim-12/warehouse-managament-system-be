package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/user/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
)

func (s *defaultUser) GetUserInfo(ctx context.Context, userId int) (resp model.UserInfoResponse, err error) {
	userData, err := s.userRepo.GetUserInfoById(ctx, userId)
	if err != nil {
		fmt.Println("failed to get user info")
		err = errors.New(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}
	resp = model.UserInfoResponse{
		UserId:       userData.ID,
		Username:     userData.Username,
		FullName:     userData.UserDetail.FullName,
		Email:        userData.Email,
		Photo:        userData.Photo,
		RegencyName:  userData.UserDetail.District.Regency.Name,
		DistrictName: userData.UserDetail.District.Name,
	}
	return
}
