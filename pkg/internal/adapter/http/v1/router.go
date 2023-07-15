package v1

import (
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
	validationService validation.Validation
	redisLibs         redis.Redis
}

func NewDateHandler(
	userService userPort.UserService,
	candidateService candidatePort.CandidateService,
	validationService validation.Validation,
	redisLibs redis.Redis) DateHandler {
	return DateHandler{
		userService:       userService,
		candidateService:  candidateService,
		validationService: validationService,
	}
}

func (handler DateHandler) RegistRoute(app *echo.Group) {
	handler.ListUserRoute(app)
	handler.ListCandidateRoute(app)
}
