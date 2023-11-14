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

func (r *defaultRepo) CreateDetail(ctx context.Context, tx *gorm.DB, req *UserDetail) (err error) {
	err = tx.WithContext(ctx).Create(req).Error
	return
}

func (r *defaultRepo) CreateUser(ctx context.Context, tx *gorm.DB, req *User) (err error) {
	err = tx.WithContext(ctx).Create(req).Error
	return
}

func (r *defaultRepo) UpdateUser(ctx context.Context, tx *gorm.DB, req *User) (err error) {
	err = tx.WithContext(ctx).Save(req).Error
	return
}

func (r *defaultRepo) BeginTrans(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx).Begin()
}

func (r *defaultRepo) DeleteUser(ctx context.Context, req *User) (err error) {
	err = r.db.WithContext(ctx).Delete(req).Error
	return
}


