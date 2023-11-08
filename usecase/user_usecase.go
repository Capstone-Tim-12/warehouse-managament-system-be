package usecase

import "github.com/Capstone-Tim-12/warehouse-managament-system-be/domain/entity"

type UserUsecase interface {
	RegisterUser(info entity.UserInfo) error
}

type userUsecase struct {
}

func NewUserUsecase() UserUsecase {
	return &userUsecase{}
}

func (u *userUsecase) RegisterUser(info entity.UserInfo) error {
	// Implement your business logic here.
	// For this example, you can just return nil.
	return nil
}
