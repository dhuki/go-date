package service

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	modelReq "github.com/dhuki/go-date/pkg/internal/adapter/http/v1/model"
	repoMock "github.com/dhuki/go-date/pkg/internal/adapter/repository/mocks"
	modelRepo "github.com/dhuki/go-date/pkg/internal/adapter/repository/model"
	redisMock "github.com/dhuki/go-date/pkg/redis/mocks"
	validationMock "github.com/dhuki/go-date/pkg/validation/mocks"
	"github.com/golang/mock/gomock"
)

func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := repoMock.NewMockRepository(ctrl)
	validationMock := validationMock.NewMockValidation(ctrl)
	redisMock := redisMock.NewMockRedis(ctrl)

	userSvc := NewUserService(userRepoMock, validationMock, redisMock)

	testCases := []struct {
		desc     string
		wantErr  error
		req      modelReq.CreateUserRequest
		mockFunc func()
	}{
		{
			desc:    "should return success sign up new user",
			wantErr: nil,
			req: modelReq.CreateUserRequest{
				Username:  "username_testing",
				Password:  "password_testing",
				FirstName: "firstname_testing",
				LastName:  "lastname_testing",
				Gender:    "gender_testing",
				PicUrl:    "pic_url_testing",
				District:  "district_testing",
				City:      "city_testing",
			},
			mockFunc: func() {
				userRepoMock.EXPECT().GetUserByUsername(gomock.Any(), gomock.Any()).Return(modelRepo.User{}, nil)
				userRepoMock.EXPECT().Start(gomock.Any()).Return(&sql.Tx{}, nil)
				userRepoMock.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any()).Return(uint64(1), nil)
				userRepoMock.EXPECT().Finish(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			desc:    "should return error when get user by username",
			wantErr: errors.New("something error"),
			req: modelReq.CreateUserRequest{
				Username:  "username_testing",
				Password:  "password_testing",
				FirstName: "firstname_testing",
				LastName:  "lastname_testing",
				Gender:    "gender_testing",
				PicUrl:    "pic_url_testing",
				District:  "district_testing",
				City:      "city_testing",
			},
			mockFunc: func() {
				userRepoMock.EXPECT().GetUserByUsername(gomock.Any(), gomock.Any()).Return(modelRepo.User{}, errors.New("something error"))
			},
		},
		{
			desc:    "should return error duplicate username",
			wantErr: errors.New("something error"),
			req: modelReq.CreateUserRequest{
				Username:  "username_testing",
				Password:  "password_testing",
				FirstName: "firstname_testing",
				LastName:  "lastname_testing",
				Gender:    "gender_testing",
				PicUrl:    "pic_url_testing",
				District:  "district_testing",
				City:      "city_testing",
			},
			mockFunc: func() {
				userRepoMock.EXPECT().GetUserByUsername(gomock.Any(), gomock.Any()).Return(modelRepo.User{
					ID: 1,
				}, nil)
			},
		},
		{
			desc:    "should return error start transaction",
			wantErr: errors.New("something error"),
			req: modelReq.CreateUserRequest{
				Username:  "username_testing",
				Password:  "password_testing",
				FirstName: "firstname_testing",
				LastName:  "lastname_testing",
				Gender:    "gender_testing",
				PicUrl:    "pic_url_testing",
				District:  "district_testing",
				City:      "city_testing",
			},
			mockFunc: func() {
				userRepoMock.EXPECT().GetUserByUsername(gomock.Any(), gomock.Any()).Return(modelRepo.User{}, nil)
				userRepoMock.EXPECT().Start(gomock.Any()).Return(&sql.Tx{}, errors.New("something error"))
			},
		},
		{
			desc:    "should return error create user",
			wantErr: errors.New("something error"),
			req: modelReq.CreateUserRequest{
				Username:  "username_testing",
				Password:  "password_testing",
				FirstName: "firstname_testing",
				LastName:  "lastname_testing",
				Gender:    "gender_testing",
				PicUrl:    "pic_url_testing",
				District:  "district_testing",
				City:      "city_testing",
			},
			mockFunc: func() {
				userRepoMock.EXPECT().GetUserByUsername(gomock.Any(), gomock.Any()).Return(modelRepo.User{}, nil)
				userRepoMock.EXPECT().Start(gomock.Any()).Return(&sql.Tx{}, nil)
				userRepoMock.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any()).Return(uint64(1), errors.New("something error"))
				userRepoMock.EXPECT().Finish(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.mockFunc != nil {
				tC.mockFunc()
			}
			err := userSvc.SignUp(context.TODO(), tC.req)

			if tC.wantErr == nil && err != nil {
				t.Fatalf("expected not error, but got error: %v", err)
			}

			if tC.wantErr != nil && err == nil {
				t.Fatalf("expected got error: %v, but actually not error", err)
			}
		})
	}
}
