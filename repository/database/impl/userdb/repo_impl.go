package userdb

import (
	"context"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/entity"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/paginate"
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

func (r *defaultRepo) GetUserInfoById(ctx context.Context, id int) (resp *entity.User, err error) {
	err = r.db.WithContext(ctx).Preload("UserDetail.District.Regency").Take(&resp, "id = ?", id).Error
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
	tx := r.db.Begin()
	err = tx.WithContext(ctx).Delete(req).Error
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.WithContext(ctx).Delete(&entity.UserDetail{}, "user_id = ?", req.ID).Error
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.WithContext(ctx).Delete(&entity.Favorit{}, "user_id = ?", req.ID).Error
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.WithContext(ctx).Delete(&entity.Transaction{}, "user_id = ?", req.ID).Error
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	return
}

func (r *defaultRepo) GetUserDetailByUserId(ctx context.Context, userId int) (resp *entity.UserDetail, err error) {
	err = r.db.WithContext(ctx).
		Preload("District").
		Preload("District.Regency").
		Preload("District.Regency.Province").
		Take(&resp, "user_id = ?", userId).
		Error
	return
}

func (r *defaultRepo) GetAllAvatar(ctx context.Context) (resp []entity.Avatar, err error) {
	err = r.db.WithContext(ctx).Find(&resp).Error
	return
}

func (r *defaultRepo) GetUserList(ctx context.Context, param paginate.Pagination) (resp []entity.User, count int64, err error) {
	query := func(condision *gorm.DB) *gorm.DB {
		if param.Search != "" {
			condision.Where("username LIKE ?", "%"+param.Search+"%")
		}
		return condision
	}
	err = r.db.WithContext(ctx).Model(&entity.User{}).Scopes(query).Count(&count).Error
	if err != nil {
		return
	}

	err = r.db.WithContext(ctx).Scopes(paginate.Paginate(param.Page, param.Limit)).Scopes(query).Order("created_at desc").Find(&resp).Error
	return
}
