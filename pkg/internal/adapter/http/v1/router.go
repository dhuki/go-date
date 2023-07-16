package v1

import (
	"github.com/dhuki/go-date/pkg/internal/adapter/repository"
	candidatePort "github.com/dhuki/go-date/pkg/internal/core/candidate/port"
	userPort "github.com/dhuki/go-date/pkg/internal/core/user/port"
	"github.com/dhuki/go-date/pkg/redis"
	"github.com/dhuki/go-date/pkg/validation"
	"github.com/labstack/echo/v4"
)

type DateHandler struct {
	userService      userPort.UserService
	candidateService candidatePort.CandidateService

	// dependency
	repository        repository.Repository
	validationService validation.Validation
	redisLibs         redis.Redis
}

func NewDateHandler(
	userService userPort.UserService,
	repository repository.Repository,
	candidateService candidatePort.CandidateService,
	validationService validation.Validation,
	redisLibs redis.Redis) DateHandler {
	return DateHandler{
		userService:       userService,
		repository:        repository,
		candidateService:  candidateService,
		validationService: validationService,
		redisLibs:         redisLibs,
	}
}

func (handler DateHandler) RegistRoute(app *echo.Group) {
	handler.ListUserRoute(app)
	handler.ListCandidateRoute(app)
}
