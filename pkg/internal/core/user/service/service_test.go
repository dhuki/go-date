package service

import (
	"context"
	"errors"
	"testing"

	"github.com/dhuki/go-date/config"
	modelReq "github.com/dhuki/go-date/pkg/internal/adapter/http/v1/model"
	repoMock "github.com/dhuki/go-date/pkg/internal/adapter/repository/mocks"
	modelRepo "github.com/dhuki/go-date/pkg/internal/adapter/repository/model"
	redisMock "github.com/dhuki/go-date/pkg/redis/mocks"
	validationMock "github.com/dhuki/go-date/pkg/validation/mocks"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
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
				userRepoMock.EXPECT().Start(gomock.Any()).Return(&sqlx.Tx{}, nil)
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
				userRepoMock.EXPECT().Start(gomock.Any()).Return(&sqlx.Tx{}, errors.New("something error"))
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
				userRepoMock.EXPECT().Start(gomock.Any()).Return(&sqlx.Tx{}, nil)
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

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := repoMock.NewMockRepository(ctrl)
	validationMock := validationMock.NewMockValidation(ctrl)
	redisMock := redisMock.NewMockRedis(ctrl)

	userSvc := NewUserService(userRepoMock, validationMock, redisMock)
	config.Conf.RateLimiter.MaxAttemptLogin = 3

	testCases := []struct {
		desc     string
		wantErr  error
		req      modelReq.LoginRequest
		mockFunc func()
	}{
		{
			desc:    "should return success login user",
			wantErr: nil,
			req: modelReq.LoginRequest{
				Username: "username_testing",
				Password: "12345",
			},
			mockFunc: func() {
				userRepoMock.EXPECT().GetUserByUsername(gomock.Any(), gomock.Any()).Return(modelRepo.User{
					ID:       1,
					Username: "username_testing",
					Password: "$2a$04$.8j.u1e7zPZ3vaXRdjnczOvio0/.Q3Wokb/H/.Up54nCdr2rx4vxa",
				}, nil)
				redisMock.EXPECT().Get(gomock.Any()).Return("")
				validationMock.EXPECT().GenerateJWTAccessToken(gomock.Any()).Return("token_testing", nil)
				redisMock.EXPECT().Delete(gomock.Any()).Return(nil)
			},
		},
		{
			desc:    "should return error get username in query",
			wantErr: errors.New("something error"),
			req: modelReq.LoginRequest{
				Username: "username_testing",
				Password: "12345",
			},
			mockFunc: func() {
				userRepoMock.EXPECT().GetUserByUsername(gomock.Any(), gomock.Any()).Return(modelRepo.User{}, errors.New("something error"))
			},
		},
		{
			desc:    "should return error get username not found",
			wantErr: errors.New("something error"),
			req: modelReq.LoginRequest{
				Username: "username_testing",
				Password: "12345",
			},
			mockFunc: func() {
				userRepoMock.EXPECT().GetUserByUsername(gomock.Any(), gomock.Any()).Return(modelRepo.User{}, nil)
			},
		},
		{
			desc:    "should return error password not match",
			wantErr: errors.New("something error"),
			req: modelReq.LoginRequest{
				Username: "username_testing",
				Password: "password_testing",
			},
			mockFunc: func() {
				userRepoMock.EXPECT().GetUserByUsername(gomock.Any(), gomock.Any()).Return(modelRepo.User{
					ID:       1,
					Username: "username_testing",
					Password: "$2a$04$.8j.u1e7zPZ3vaXRdjnczOvio0/.Q3Wokb/H/.Up54nCdr2rx4vxa",
				}, nil)
				redisMock.EXPECT().Get(gomock.Any()).Return("")
				redisMock.EXPECT().SetIncr(gomock.Any()).Return(int64(1))
			},
		},
		{
			desc:    "should return error password not match then suspend",
			wantErr: errors.New("something error"),
			req: modelReq.LoginRequest{
				Username: "username_testing",
				Password: "password_testing",
			},
			mockFunc: func() {
				userRepoMock.EXPECT().GetUserByUsername(gomock.Any(), gomock.Any()).Return(modelRepo.User{
					ID:       1,
					Username: "username_testing",
					Password: "$2a$04$.8j.u1e7zPZ3vaXRdjnczOvio0/.Q3Wokb/H/.Up54nCdr2rx4vxa",
				}, nil)
				redisMock.EXPECT().Get(gomock.Any()).Return("")
				redisMock.EXPECT().SetIncr(gomock.Any()).Return(int64(3))
				redisMock.EXPECT().Delete(gomock.Any()).Return(nil)
				redisMock.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			desc:    "should return error suspend",
			wantErr: errors.New("something error"),
			req: modelReq.LoginRequest{
				Username: "username_testing",
				Password: "password_testing",
			},
			mockFunc: func() {
				userRepoMock.EXPECT().GetUserByUsername(gomock.Any(), gomock.Any()).Return(modelRepo.User{
					ID:       1,
					Username: "username_testing",
					Password: "$2a$04$.8j.u1e7zPZ3vaXRdjnczOvio0/.Q3Wokb/H/.Up54nCdr2rx4vxa",
				}, nil)
				redisMock.EXPECT().Get(gomock.Any()).Return("true")
			},
		},
		{
			desc:    "should return error generate access token",
			wantErr: errors.New("something error"),
			req: modelReq.LoginRequest{
				Username: "username_testing",
				Password: "12345",
			},
			mockFunc: func() {
				userRepoMock.EXPECT().GetUserByUsername(gomock.Any(), gomock.Any()).Return(modelRepo.User{
					ID:       1,
					Username: "username_testing",
					Password: "$2a$04$.8j.u1e7zPZ3vaXRdjnczOvio0/.Q3Wokb/H/.Up54nCdr2rx4vxa",
				}, nil)
				redisMock.EXPECT().Get(gomock.Any()).Return("")
				validationMock.EXPECT().GenerateJWTAccessToken(gomock.Any()).Return("", errors.New("something error"))
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.mockFunc != nil {
				tC.mockFunc()
			}
			_, err := userSvc.Login(context.TODO(), tC.req)

			if tC.wantErr == nil && err != nil {
				t.Fatalf("expected not error, but got error: %v", err)
			}

			if tC.wantErr != nil && err == nil {
				t.Fatalf("expected got error: %v, but actually not error", err)
			}
		})
	}
}
