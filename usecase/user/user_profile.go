package user

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/user/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/errors"
	"github.com/spf13/cast"
)

func (s *defaultUser) UpdateUsernameProfile(ctx context.Context, userId string, req model.UpdateUsernameProfileRequest) (err error) {
	userData, err := s.userRepo.GetUserById(ctx, cast.ToInt(userId))
	if err != nil {
		fmt.Println("user not found")
		err = errors.New(http.StatusNotFound, "user not found")
		return
	}

	user, _ := s.userRepo.GetUserByUsername(ctx, req.Username)
	if user.Username != "" {
		fmt.Println("username already set")
		err = errors.New(http.StatusConflict, "username already set in other user")
		return
	}

	userData.Username = req.Username
	tx := s.userRepo.BeginTrans(ctx)
	err = s.userRepo.UpdateUser(ctx, tx, userData)
	if err != nil {
		tx.Rollback()
		fmt.Println("error update user: ", err.Error())
		err = errors.New(http.StatusInternalServerError, "error updating user")
		return
	}

	tx.Commit()
	return
}

func (s *defaultUser) UpdatePhotoProfile(ctx context.Context, userId int, image *multipart.FileHeader) (err error) {
	userData, err := s.userRepo.GetUserById(ctx, cast.ToInt(userId))
	if err != nil {
		fmt.Println("user not found")
		err = errors.New(http.StatusNotFound, "user not found")
		return
	}

	urlImage, err := s.coreRepo.UploadImage(ctx, image)
	if err != nil {
		fmt.Println("error uploading image: ", err.Error())
		err = errors.New(http.StatusInternalServerError, "error uploading image")
		return
	}

	if len(urlImage.Data.Images) == 0 {
		fmt.Println("failed upload images")
		err = errors.New(http.StatusInternalServerError, "error uploading image")
		return
	}
	tx := s.userRepo.BeginTrans(ctx)
	userData.Photo =  urlImage.Data.Images[0]
	err = s.userRepo.UpdateUser(ctx, tx, userData)
	if err != nil {
		tx.Rollback()
		fmt.Println("error updating user")
		err = errors.New(http.StatusInternalServerError, "error updating user")
		return
	}

	tx.Commit()
	return
}

func (s *defaultUser) GetProfile(ctx context.Context, userId string) (resp model.GetProfileResponse, err error) {
	userData, err := s.userRepo.GetUserById(ctx, cast.ToInt(userId))
	if err != nil {
		fmt.Println("user not found")
		err = errors.New(http.StatusNotFound, "user not found")
		return
	}

	resp = model.GetProfileResponse{
		Id:               userData.ID,
		Username:         userData.Username,
		Email:            userData.Email,
		IsVerifyAccount:  userData.IsVerifyAccount,
		IsVerifyIdentity: userData.IsVerifyIdentity,
		Photo:            userData.Photo,
	}

	userDetail, _ := s.userRepo.GetUserDetailByUserId(ctx, userData.ID)
	if userDetail.ID != 0 {
		resp.Address = userDetail.Address
		resp.Country = userDetail.Country
		resp.NIK = userDetail.NIK
		resp.FullName = userDetail.FullName
		resp.Gender = userDetail.Gender
		resp.PlaceOfBirth = userDetail.PlaceOfBirth
		resp.DateBirth = userDetail.DateBirth
		resp.Work = userDetail.Work
		resp.Citizenship = userDetail.Citizenship
		resp.ProvinceID = userDetail.ProvinceID
		resp.ProvinceName = userDetail.Province.Name
		resp.RegencyID = userDetail.RegencyID
		resp.RegencyName = userDetail.Regency.Name
		resp.DistrictID = userDetail.DistrictID
		resp.DistrictName = userDetail.District.Name
	}

	return
}