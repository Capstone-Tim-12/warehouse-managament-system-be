package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
)

func (s *defaultUser) DeleteUser(ctx context.Context, id int) (err error) {
	userData, err := s.userRepo.GetUserById(ctx, id)
	if err != nil {
		fmt.Println("Error getting user by Id", err.Error())
		err = errors.New(http.StatusNotFound, "user not found")
		return
	}
	err = s.userRepo.DeleteUser(ctx, userData)
	if err != nil {
		fmt.Println("Error delete user", err.Error())
		err = errors.New(http.StatusInternalServerError, "Can't delete user")
		return
	}
	return
}
