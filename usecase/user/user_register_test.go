package user

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/database/entity"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/repository/http/core"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/usecase/user/model"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils"
	"github.com/Capstone-Tim-12/warehouse-managament-system-be/utils/mocks"
	"github.com/stretchr/testify/mock"
)

func Test_defaultUser_UserRegister(t *testing.T) {
	db := utils.MockGorm()
	ctx := context.TODO()

	type args struct {
		ctx  context.Context
		req  model.RegisterUserRequest
		long float64
		lat  float64
	}
	tests := []struct {
		name           string
		args           args
		getByEmailResp *entity.User
		getByEmailErr  error
		createUserErr  error
		setUtilityResp core.SetUtilityResponse
		setUtilityErr  error
		sendEmailResp  core.SendEmailResponse
		sendEmailErr   error
		wantResp       model.RegisterUserResponse
		wantErr        bool
	}{
		{
			name: "email already registered",
			args: args{
				ctx: ctx,
				req: model.RegisterUserRequest{
					Username: "fadilah",
					Email:    "fadilah@gmail.com",
					Password: "123456",
				},
				long: 15.87,
				lat:  19.67,
			},
			getByEmailResp: &entity.User{
				Email: "fadilah@gmail.com",
			},
			wantErr: true,
		},
		{
			name: "error create user",
			args: args{
				ctx: ctx,
				req: model.RegisterUserRequest{
					Username: "fadilah",
					Email:    "fadilah@gmail.com",
					Password: "123456",
				},
				long: 15.87,
				lat:  19.67,
			},
			getByEmailResp: &entity.User{
				Email: "",
			},
			createUserErr: errors.New("error creating user"),
			wantErr: true,
		},
		{
			name: "set utility error",
			args: args{
				ctx: ctx,
				req: model.RegisterUserRequest{
					Username: "fadilah",
					Email:    "fadilah@gmail.com",
					Password: "123456",
				},
				long: 15.87,
				lat:  19.67,
			},
			getByEmailResp: &entity.User{
				Email: "",
			},
			setUtilityErr: errors.New("set utility error"),
			wantErr: true,
		},
		{
			name: "send email failed error",
			args: args{
				ctx: ctx,
				req: model.RegisterUserRequest{
					Username: "fadilah",
					Email:    "fadilah@gmail.com",
					Password: "123456",
				},
				long: 15.87,
				lat:  19.67,
			},
			getByEmailResp: &entity.User{
				Email: "",
			},
			sendEmailErr: errors.New("send email error"),
			wantErr: true,
		},
		{
			name: "register user success",
			args: args{
				ctx: ctx,
				req: model.RegisterUserRequest{
					Username: "fadilah",
					Email:    "fadilah@gmail.com",
					Password: "123456",
				},
				long: 15.87,
				lat:  19.67,
			},
			getByEmailResp: &entity.User{
				Email: "",
			},
			wantResp: model.RegisterUserResponse{
				Email: "fadilah@gmail.com",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepo := new(mocks.UserRepository)
			regionRepo := new(mocks.RegionRepository)
			coreWrapper := new(mocks.CoreWrapper)
			chatWrapper := new(mocks.OpenAIWrapper)

			userRepo.On("GetUserByEmail", mock.Anything, mock.Anything).Return(tt.getByEmailResp, tt.getByEmailErr).Once()
			userRepo.On("BeginTrans", mock.Anything).Return(db).Once()
			userRepo.On("CreateUser", mock.Anything, mock.Anything, mock.Anything).Return(tt.createUserErr).Once()
			coreWrapper.On("SetUtility", mock.Anything, mock.Anything).Return(tt.setUtilityResp, tt.setUtilityErr).Once()
			coreWrapper.On("SendEmail", mock.Anything, mock.Anything).Return(tt.sendEmailResp, tt.sendEmailErr).Once()

			svc := NewUserUsecase(regionRepo, userRepo, coreWrapper, chatWrapper)
			gotResp, err := svc.UserRegister(tt.args.ctx, tt.args.req, tt.args.long, tt.args.lat)
			if (err != nil) != tt.wantErr {
				t.Errorf("defaultUser.UserRegister() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("defaultUser.UserRegister() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
