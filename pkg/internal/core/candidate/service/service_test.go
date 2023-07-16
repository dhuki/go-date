package service

import (
	"context"
	"errors"
	"testing"

	"github.com/dhuki/go-date/config"
	repoMock "github.com/dhuki/go-date/pkg/internal/adapter/repository/mocks"
	modelRepo "github.com/dhuki/go-date/pkg/internal/adapter/repository/model"
	redisMock "github.com/dhuki/go-date/pkg/redis/mocks"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
)

func TestSwipeAction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	candidateRepoMock := repoMock.NewMockRepository(ctrl)
	redisMock := redisMock.NewMockRedis(ctrl)
	candidateSvc := NewCandidateService(candidateRepoMock, redisMock)

	ctx := context.WithValue(context.TODO(), config.ValueUserIDctx, "1")

	testCases := []struct {
		desc           string
		candidateID    uint64
		swipeDirection string
		wantErr        error
		mockFunc       func()
	}{
		{
			desc:           "should return succes when swipe action",
			candidateID:    1,
			swipeDirection: "right",
			wantErr:        nil,
			mockFunc: func() {
				candidateRepoMock.EXPECT().GetRelationUserByUserIdAndCandidate(gomock.Any(), gomock.Any(), gomock.Any()).Return(modelRepo.RelationUser{}, nil)
				candidateRepoMock.EXPECT().Start(gomock.Any()).Return(&sqlx.Tx{}, nil)
				candidateRepoMock.EXPECT().UpsertRelationUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
				candidateRepoMock.EXPECT().Finish(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			desc:           "should return error invalid swipe action",
			candidateID:    1,
			swipeDirection: "invalid_direction",
			wantErr:        errors.New("something error"),
			mockFunc:       nil,
		},
		{
			desc:           "should return error get relation user",
			candidateID:    1,
			swipeDirection: "right",
			wantErr:        errors.New("something error"),
			mockFunc: func() {
				candidateRepoMock.EXPECT().GetRelationUserByUserIdAndCandidate(gomock.Any(), gomock.Any(), gomock.Any()).Return(modelRepo.RelationUser{}, errors.New("something error"))
			},
		},
		{
			desc:           "should return error when upsert relation user",
			candidateID:    1,
			swipeDirection: "right",
			wantErr:        errors.New("something error"),
			mockFunc: func() {
				candidateRepoMock.EXPECT().GetRelationUserByUserIdAndCandidate(gomock.Any(), gomock.Any(), gomock.Any()).Return(modelRepo.RelationUser{}, nil)
				candidateRepoMock.EXPECT().Start(gomock.Any()).Return(&sqlx.Tx{}, nil)
				candidateRepoMock.EXPECT().UpsertRelationUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("something error"))
				candidateRepoMock.EXPECT().Finish(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.mockFunc != nil {
				tC.mockFunc()
			}
			err := candidateSvc.SwipeAction(ctx, tC.candidateID, tC.swipeDirection)

			if tC.wantErr == nil && err != nil {
				t.Fatalf("expected not error, but got error: %v", err)
			}

			if tC.wantErr != nil && err == nil {
				t.Fatalf("expected got error: %v, but actually not error", err)
			}
		})
	}
}

func TestGetListCandidate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	candidateRepoMock := repoMock.NewMockRepository(ctrl)
	redisMock := redisMock.NewMockRedis(ctrl)
	candidateSvc := NewCandidateService(candidateRepoMock, redisMock)

	ctx := context.WithValue(context.TODO(), config.ValueUserIDctx, "1")

	testCases := []struct {
		desc     string
		limit    int
		wantErr  error
		mockFunc func()
	}{
		{
			desc:    "should return succes when get empty list candidate",
			limit:   0,
			wantErr: nil,
			mockFunc: func() {
				candidateRepoMock.EXPECT().GetUserByID(gomock.Any(), gomock.Any()).Return(modelRepo.User{
					ID:        1,
					Username:  "username_testing",
					Password:  "password_testing",
					FirstName: "firstname_testing",
					LastName:  "lastname_testing",
					Gender:    "gender_testing",
					PicUrl:    "pic_url_testing",
					District:  "district_testing",
					City:      "city_testing",
				}, nil)
				candidateRepoMock.EXPECT().CountTotalUserPagination(gomock.Any(), gomock.Any(), gomock.Any()).Return(int(1), nil)
			},
		},
		{
			desc:    "should return succes when get list candidate",
			limit:   1,
			wantErr: nil,
			mockFunc: func() {
				candidateRepoMock.EXPECT().GetUserByID(gomock.Any(), gomock.Any()).Return(modelRepo.User{
					ID:        1,
					Username:  "username_testing",
					Password:  "password_testing",
					FirstName: "firstname_testing",
					LastName:  "lastname_testing",
					Gender:    "gender_testing",
					PicUrl:    "pic_url_testing",
					District:  "district_testing",
					City:      "city_testing",
				}, nil)
				candidateRepoMock.EXPECT().CountTotalUserPagination(gomock.Any(), gomock.Any(), gomock.Any()).Return(int(1), nil)
				redisMock.EXPECT().Get(gomock.Any()).Return("")
				candidateRepoMock.EXPECT().GetUserPagination(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]modelRepo.User{
					{
						ID:        2,
						Username:  "username_testing",
						Password:  "password_testing",
						FirstName: "firstname_testing",
						LastName:  "lastname_testing",
						Gender:    "gender_testing",
						PicUrl:    "pic_url_testing",
						District:  "district_testing",
						City:      "city_testing",
					},
				}, nil)
				redisMock.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			desc:    "should return succes when get list candidate by previous page",
			limit:   1,
			wantErr: nil,
			mockFunc: func() {
				candidateRepoMock.EXPECT().GetUserByID(gomock.Any(), gomock.Any()).Return(modelRepo.User{
					ID:        1,
					Username:  "username_testing",
					Password:  "password_testing",
					FirstName: "firstname_testing",
					LastName:  "lastname_testing",
					Gender:    "gender_testing",
					PicUrl:    "pic_url_testing",
					District:  "district_testing",
					City:      "city_testing",
				}, nil)
				candidateRepoMock.EXPECT().CountTotalUserPagination(gomock.Any(), gomock.Any(), gomock.Any()).Return(int(2), nil)
				redisMock.EXPECT().Get(gomock.Any()).Return("1")
				candidateRepoMock.EXPECT().GetUserPagination(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]modelRepo.User{
					{
						ID:        2,
						Username:  "username_testing",
						Password:  "password_testing",
						FirstName: "firstname_testing",
						LastName:  "lastname_testing",
						Gender:    "gender_testing",
						PicUrl:    "pic_url_testing",
						District:  "district_testing",
						City:      "city_testing",
					},
				}, nil)
				redisMock.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			desc:    "should return error get user by id",
			limit:   0,
			wantErr: errors.New("something error"),
			mockFunc: func() {
				candidateRepoMock.EXPECT().GetUserByID(gomock.Any(), gomock.Any()).Return(modelRepo.User{}, errors.New("something error"))
			},
		},
		{
			desc:    "should return error get count total user",
			limit:   0,
			wantErr: errors.New("something error"),
			mockFunc: func() {
				candidateRepoMock.EXPECT().GetUserByID(gomock.Any(), gomock.Any()).Return(modelRepo.User{}, nil)
				candidateRepoMock.EXPECT().CountTotalUserPagination(gomock.Any(), gomock.Any(), gomock.Any()).Return(0, errors.New("something error"))
			},
		},
		{
			desc:    "should return error get list data user",
			limit:   1,
			wantErr: errors.New("something error"),
			mockFunc: func() {
				candidateRepoMock.EXPECT().GetUserByID(gomock.Any(), gomock.Any()).Return(modelRepo.User{
					ID:        1,
					Username:  "username_testing",
					Password:  "password_testing",
					FirstName: "firstname_testing",
					LastName:  "lastname_testing",
					Gender:    "gender_testing",
					PicUrl:    "pic_url_testing",
					District:  "district_testing",
					City:      "city_testing",
				}, nil)
				candidateRepoMock.EXPECT().CountTotalUserPagination(gomock.Any(), gomock.Any(), gomock.Any()).Return(1, nil)
				redisMock.EXPECT().Get(gomock.Any()).Return("")
				candidateRepoMock.EXPECT().GetUserPagination(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]modelRepo.User{}, errors.New("something error"))
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.mockFunc != nil {
				tC.mockFunc()
			}
			_, err := candidateSvc.GetListCandidate(ctx, tC.limit)

			if tC.wantErr == nil && err != nil {
				t.Fatalf("expected not error, but got error: %v", err)
			}

			if tC.wantErr != nil && err == nil {
				t.Fatalf("expected got error: %v, but actually not error", err)
			}
		})
	}
}
