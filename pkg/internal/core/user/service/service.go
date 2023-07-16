package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/dhuki/go-date/config"
	modelReq "github.com/dhuki/go-date/pkg/internal/adapter/http/v1/model"
	modelRepo "github.com/dhuki/go-date/pkg/internal/adapter/repository/model"
	"github.com/dhuki/go-date/pkg/internal/core/user/domain"
	"github.com/dhuki/go-date/pkg/internal/core/user/port"
	"github.com/dhuki/go-date/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

type userServiceImpl struct {
	repository port.UserRepository
	validation port.JWTAccessToken
	redisLib   port.RedisLibs
}

func NewUserService(userRepository port.UserRepository, validation port.JWTAccessToken, redisLib port.RedisLibs) port.UserService {
	return userServiceImpl{
		repository: userRepository,
		validation: validation,
		redisLib:   redisLib,
	}
}

func (u userServiceImpl) SignUp(ctx context.Context, req modelReq.CreateUserRequest) (err error) {
	ctxName := fmt.Sprintf("%T.SignUp", u)

	user, err := u.repository.GetUserByUsername(ctx, req.Username)
	if err != nil && err != sql.ErrNoRows {
		logger.Error(ctx, ctxName, "u.repository.GetUserByUsername, got err: %v", err)
		return
	}

	if user.ID > 0 {
		err = domain.ErrUserAlreadyExist
		logger.Error(ctx, ctxName, "found.username, got err: %v", err)
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
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Gender:    req.Gender,
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

	if user.ID <= 0 {
		err = domain.ErrUsernameIsNotFound
		return
	}

	keyFailedPassword := fmt.Sprintf("%s.%d", domain.KeyMismatchPassword, user.ID)
	keyTemporarySuspend := fmt.Sprintf("%s.%d", domain.KeyTemporarySuspend, user.ID)

	if value := u.redisLib.Get(keyTemporarySuspend); len(value) > 0 {
		err = domain.ErrTooManyFailedLogin
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			failedAttempt := u.redisLib.SetIncr(keyFailedPassword)
			if int(failedAttempt) >= config.Conf.RateLimiter.MaxAttemptLogin {
				u.temporarySuspendAccount(ctx, keyFailedPassword, keyTemporarySuspend)
				err = domain.ErrTooManyFailedLogin
				return
			}
			err = domain.ErrWrongPassword
			return
		}
		logger.Error(ctx, ctxName, "bcrypt.CompareHashAndPassword, got err: %v", err)
		return
	}

	resp.AccessToken, err = u.validation.GenerateJWTAccessToken(user.ID)
	if err != nil {
		logger.Error(ctx, ctxName, "u.validation.GenerateJWTAccessToken, got err: %v", err)
		return
	}

	resp.Username = user.Username
	resp.FirstName = user.FirstName
	resp.LastName = user.LastName
	resp.Gender = user.Gender
	resp.PicUrl = user.PicUrl
	resp.District = user.District
	resp.City = user.City

	if err = u.redisLib.Delete(keyFailedPassword); err != nil {
		logger.Warn(ctx, ctxName, "u.redisLib.Delete, got err: %v", err)
	}

	return
}
