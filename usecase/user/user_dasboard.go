package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/user/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/paginate"
)

func (s *defaultUser) GetUserList(ctx context.Context, param paginate.Pagination) (resp []model.UserListResponse, count int64, err error) {
	userData, count, err := s.userRepo.GetUserList(ctx, param)
	if err != nil {
		fmt.Println("error getting user list: ", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	for i := 0; i < len(userData); i++ {
		resp = append(resp, model.UserListResponse{
			UserId:           userData[i].ID,
			Username:         userData[i].Username,
			Email:            userData[i].Email,
			Photo:            userData[i].Photo,
			IsVerifyIdentity: userData[i].IsVerifyIdentity,
		})
	}
	return
}

func (s *defaultUser) GetUserById(ctx context.Context, userId int) (resp model.UserListResponse, err error) {
	userData, err := s.userRepo.GetUserById(ctx, userId)
	if err != nil {
		fmt.Println("user not found: ", err.Error())
		err = errors.New(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}
	resp = model.UserListResponse{
		UserId:   userData.ID,
		Username: userData.Username,
		Email:    userData.Email,
		Photo:    userData.Photo,
	}
	return
}
