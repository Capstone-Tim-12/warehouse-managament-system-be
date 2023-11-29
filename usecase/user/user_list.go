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
			UserId:   userData[i].ID,
			Username: userData[i].Username,
			Email:    userData[i].Email,
			Photo:    userData[i].Photo,
		})
	}
	return
}