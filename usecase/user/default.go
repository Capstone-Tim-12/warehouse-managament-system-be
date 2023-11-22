package user

import (
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/impl/regiondb"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/impl/userdb"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/http/core"
)

type defaultUser struct {
	regionRepo regiondb.RegionRepository
	userRepo   userdb.UserRepository
	coreRepo   core.CoreWrapper
}

func NewUserUsecase(regionRepo regiondb.RegionRepository, userRepo userdb.UserRepository, coreRepo core.CoreWrapper) UserUsecase {
	return &defaultUser{
		regionRepo: regionRepo,
		userRepo:   userRepo,
		coreRepo:   coreRepo,
	}
}
