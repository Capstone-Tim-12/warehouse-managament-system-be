package userdb

import (
	"context"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/entity"
	"gorm.io/gorm"
)

type defaultRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &defaultRepo{db: db}
}

func (r *defaultRepo) GetUserByEmail(ctx context.Context, email string) (resp *entity.User, err error) {
	err = r.db.WithContext(ctx).Take(&resp, "email = ?", email).Error
	return
}

func (r *defaultRepo) GetUserById(ctx context.Context, id int) (resp *entity.User, err error) {
	err = r.db.WithContext(ctx).Take(&resp, "id = ?", id).Error
	return
}

func (r *defaultRepo) GetUserByUsername(ctx context.Context, username string) (resp *entity.User, err error) {
	err = r.db.WithContext(ctx).Take(&resp, "username = ?", username).Error
	return
}

func (r *defaultRepo) GetUserByEmailUsername(ctx context.Context, email, username string) (resp *entity.User, err error) {
	err = r.db.WithContext(ctx).Take(&resp, "email = ? AND username = ?", email, username).Error
	return
}

func (r *defaultRepo) CreateDetail(ctx context.Context, tx *gorm.DB, req *entity.UserDetail) (err error) {
	err = tx.WithContext(ctx).Create(req).Error
	return
}

func (r *defaultRepo) CreateUser(ctx context.Context, tx *gorm.DB, req *entity.User) (err error) {
	err = tx.WithContext(ctx).Create(req).Error
	return
}

func (r *defaultRepo) UpdateUser(ctx context.Context, tx *gorm.DB, req *entity.User) (err error) {
	err = tx.WithContext(ctx).Save(req).Error
	return
}

func (r *defaultRepo) BeginTrans(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx).Begin()
}

func (r *defaultRepo) DeleteUser(ctx context.Context, req *entity.User) (err error) {
	err = r.db.WithContext(ctx).Delete(req).Error
	return
}

func (r *defaultRepo) GetUserDetailByUserId(ctx context.Context, userId int) (resp *entity.UserDetail, err error) {
	err = r.db.WithContext(ctx).Preload("Province").
		Preload("Regency").
		Preload("District").
		Take(&resp, "user_id = ?", userId).
		Error
	return
}

func (r *defaultRepo) GetAllAvatar(ctx context.Context) (resp []entity.Avatar, err error) {
	err = r.db.WithContext(ctx).Find(&resp).Error
	return
}


