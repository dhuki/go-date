package service

import (
	"context"
	"fmt"
	"time"

	modelReq "github.com/dhuki/go-date/pkg/internal/adapter/http/v1/model"
	modelRepo "github.com/dhuki/go-date/pkg/internal/adapter/repository/model"
	"github.com/dhuki/go-date/pkg/internal/core/user/port"
	"github.com/dhuki/go-date/pkg/logger"
	validation "github.com/dhuki/go-date/pkg/validation"
	"golang.org/x/crypto/bcrypt"
)

type userServiceImpl struct {
	repository port.UserRepository
	validation validation.Validation
}

func NewUserService(userRepository port.UserRepository, validation validation.Validation) port.UserService {
	return userServiceImpl{
		repository: userRepository,
		validation: validation,
	}
}

func (u userServiceImpl) SignUp(ctx context.Context, req modelReq.CreateUserRequest) (err error) {
	ctxName := fmt.Sprintf("%T.SignUp", u)

	user, err := u.repository.GetUserByUsername(ctx, req.Username)
	if err != nil {
		logger.Error(ctx, ctxName, "u.repository.GetUserByUsername, got err: %v", err)
		return
	}

	if user.ID > 0 {
		return
	}

	encryptPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	req.Password = string(encryptPass)

	tx, err := u.repository.Start(ctx)
	if err != nil {
		logger.Error(ctx, ctxName, "u.repository.Start, got err: %v", err)
		return
	}
	defer func() {
		if err := u.repository.Finish(ctx, tx, err); err != nil {
			logger.Error(ctx, ctxName, "u.repository.Finish, err: %v", err)
		}
	}()

	_, err = u.repository.Create(ctx, tx, modelRepo.User{
		Username:  req.Username,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		PicUrl:    req.PicUrl,
		District:  req.District,
		City:      req.City,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		logger.Error(ctx, ctxName, "u.repository.Create, got err: %v", err)
		return
	}
	return
}

func (u userServiceImpl) Login(ctx context.Context, req modelReq.LoginRequest) (resp modelReq.LoginResponse, err error) {
	ctxName := fmt.Sprintf("%T.Login", u)

	user, err := u.repository.GetUserByUsername(ctx, req.Username)
	if err != nil {
		logger.Error(ctx, ctxName, "u.repository.GetUserByUsername, got err: %v", err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(req.Password), []byte(req.Password))
	if err != nil {
		logger.Error(ctx, ctxName, "bcrypt.CompareHashAndPassword, got err: %v", err)
		return
	}

	resp.AccessToken, err = u.validation.GenerateJWTAccessToken()
	if err != nil {
		logger.Error(ctx, ctxName, "u.validation.GenerateJWTAccessToken, got err: %v", err)
		return
	}

	resp.Username = user.Username
	resp.FirstName = user.FirstName
	resp.LastName = user.LastName
	resp.Email = user.Email
	resp.PicUrl = user.PicUrl
	resp.District = user.District
	resp.City = user.City

	return
}
