package userdb

import (
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (resp *User, err error)
	CreateDetail(ctx context.Context, tx *gorm.DB, req *UserDetail) (err error)
	CreateUser(ctx context.Context, tx *gorm.DB, req *User) (err error)
	UpdateUser(ctx context.Context, tx *gorm.DB, req *User) (err error)
	BeginTrans(ctx context.Context) *gorm.DB 
	DeleteUser(ctx context.Context, req *User) (err error)
}
