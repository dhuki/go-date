package service

import (
	"context"
	"testing"

	"github.com/dhuki/go-date/pkg/internal/adapter/http/v1/model"
	repoMock "github.com/dhuki/go-date/pkg/internal/adapter/repository/mocks"
	validationMock "github.com/dhuki/go-date/pkg/validation/mocks"
	"github.com/golang/mock/gomock"
)

func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := repoMock.NewMockRepository(ctrl)
	validationMock := validationMock.NewMockValidation(ctrl)

	userSvc := NewUserService(userRepoMock, validationMock)

	testCases := []struct {
		desc string
		req  model.CreateUserRequest
	}{
		{
			desc: "should return success sign up new user",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			userSvc.SignUp(context.TODO(), tC.req)
		})
	}
}
