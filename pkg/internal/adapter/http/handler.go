package http

import (
	"github.com/dhuki/go-date/config/database"
	v1 "github.com/dhuki/go-date/pkg/internal/adapter/http/v1"
	"github.com/dhuki/go-date/pkg/internal/adapter/repository"
	candidateService "github.com/dhuki/go-date/pkg/internal/core/candidate/service"
	userService "github.com/dhuki/go-date/pkg/internal/core/user/service"
	"github.com/dhuki/go-date/pkg/middleware"
	"github.com/dhuki/go-date/pkg/validation"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func NewHttpHandler(app *echo.Echo) {
	app.Use(echoMiddleware.TimeoutWithConfig(echoMiddleware.TimeoutConfig{
		Timeout:      60,
		ErrorMessage: "timeout",
	}))
	app.Use(echoMiddleware.Recover())
	app.Use(middleware.Logger())

	v1Group := app.Group("/api/v1")

	// repository
	repo := repository.NewRepository(database.PostgresDb.Master, database.PostgresDb.Slave)

	// user service
	validationSvc := validation.NewValidation()
	userSvc := userService.NewUserService(repo, validationSvc)

	// candidate service
	repoSvc := candidateService.NewCandidateService(repo)

	v1.NewDateHandler(userSvc, repoSvc).RegistRoute(v1Group)
}
