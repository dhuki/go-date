package v1

import (
	candidatePort "github.com/dhuki/go-date/pkg/internal/core/candidate/port"
	userPort "github.com/dhuki/go-date/pkg/internal/core/user/port"
	"github.com/labstack/echo/v4"
)

type DateHandler struct {
	userService      userPort.UserService
	candidateService candidatePort.CandidateService
}

func NewDateHandler(userService userPort.UserService, candidateService candidatePort.CandidateService) DateHandler {
	return DateHandler{
		userService:      userService,
		candidateService: candidateService,
	}
}

func (handler DateHandler) RegistRoute(app *echo.Group) {
	handler.ListUserRoute(app)
	handler.ListCandidateRoute(app)
}
