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

func (s *defaultUser) UpdatePhotoProfile(ctx context.Context, userId int, req model.UpdatePhotoProfileRequest) (err error) {
	userData, err := s.userRepo.GetUserById(ctx, cast.ToInt(userId))
	if err != nil {
		fmt.Println("user not found")
		err = errors.New(http.StatusNotFound, "user not found")
		return
	}

	tx := s.userRepo.BeginTrans(ctx)
	userData.Photo = req.UrlImage
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
		resp.ProvinceID = userDetail.District.Regency.ProvinceID
		resp.ProvinceName = userDetail.District.Regency.Province.Name
		resp.RegencyID = userDetail.District.RegencyID
		resp.RegencyName = userDetail.District.Regency.Name
		resp.DistrictID = userDetail.DistrictID
		resp.DistrictName = userDetail.District.Name
	}

	return
}

func (s *defaultUser) UploadPhoto(ctx context.Context, image *multipart.FileHeader) (resp model.UploadPhotoResponse, err error) {
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

	resp.UrlImage = urlImage.Data.Images[0]
	return
}

func (s *defaultUser) GetAvatarList(ctx context.Context) (resp []model.GetAvatarResponse, err error) {
	data, err := s.userRepo.GetAllAvatar(ctx)
	if err != nil {
		fmt.Println("error getting avatar list: ", err.Error())
		err = errors.New(http.StatusInternalServerError, "error getting avatar list")
		return
	}
	for i := 0; i < len(data); i++ {
		resp = append(resp, model.GetAvatarResponse{
			Id:    data[i].ID,
			Image: data[i].Image,
		})
	}
	return
}

func (s *defaultUser) UpdateEmail(ctx context.Context, userId int, req model.OtpRequest) (err error) {
	userData, _ := s.userRepo.GetUserByEmail(ctx, req.Email)
	if userData.Email != "" {
		fmt.Println("email is already existing")
		err = errors.New(http.StatusInternalServerError, "email is already existing")
		return
	}
	userData, err = s.userRepo.GetUserById(ctx, userId)
	if err != nil {
		fmt.Println("user not found")
		err = errors.New(http.StatusNotFound, "user not found")
		return
	}

	userData.Email = req.Email
	userData.IsVerifyAccount = false
	tx := s.userRepo.BeginTrans(ctx)
	err = s.userRepo.UpdateUser(ctx, tx, userData)
	if err != nil {
		tx.Rollback()
		fmt.Println("failed to update user")
		err = errors.New(http.StatusInternalServerError, "failed to update email")
		return
	}

	err = s.sendEmailOtp(ctx, *userData)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	return
}