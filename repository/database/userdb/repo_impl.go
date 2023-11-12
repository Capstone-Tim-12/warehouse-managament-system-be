package userdb

import (
	"context"

	"gorm.io/gorm"
)

type defaultRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &defaultRepo{db: db}
}

func (r *defaultRepo) GetUserByEmail(ctx context.Context, email string) (resp *User, err error) {
	err = r.db.WithContext(ctx).Take(&resp, "email = ?", email).Error
	return
}

func (r *defaultRepo) Create(ctx context.Context, req *UserDetail) (err error) {
	err = r.db.WithContext(ctx).Create(req).Error
	return
}
