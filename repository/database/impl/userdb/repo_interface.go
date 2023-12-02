package userdb

import (
	"context"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/entity"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/paginate"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (resp *entity.User, err error)
	CreateDetail(ctx context.Context, tx *gorm.DB, req *entity.UserDetail) (err error)
	CreateUser(ctx context.Context, tx *gorm.DB, req *entity.User) (err error)
	UpdateUser(ctx context.Context, tx *gorm.DB, req *entity.User) (err error)
	BeginTrans(ctx context.Context) *gorm.DB
	DeleteUser(ctx context.Context, req *entity.User) (err error)
	GetUserByUsername(ctx context.Context, username string) (resp *entity.User, err error)
	GetUserByEmailUsername(ctx context.Context, email, username string) (resp *entity.User, err error)
	GetUserById(ctx context.Context, id int) (resp *entity.User, err error)
	GetUserDetailByUserId(ctx context.Context, userId int) (resp *entity.UserDetail, err error)
	GetAllAvatar(ctx context.Context) (resp []entity.Avatar, err error)
	GetUserList(ctx context.Context, param paginate.Pagination) (resp []entity.User, count int64, err error)
	GetUserInfoById(ctx context.Context, id int) (resp *entity.User, err error) 
}
