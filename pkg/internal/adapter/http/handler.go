package http

import (
	"github.com/dhuki/go-date/config"
	"github.com/dhuki/go-date/config/database"
	redisClient "github.com/dhuki/go-date/config/redis"
	v1 "github.com/dhuki/go-date/pkg/internal/adapter/http/v1"
	"github.com/dhuki/go-date/pkg/internal/adapter/repository"
	candidateService "github.com/dhuki/go-date/pkg/internal/core/candidate/service"
	userService "github.com/dhuki/go-date/pkg/internal/core/user/service"
	"github.com/dhuki/go-date/pkg/middleware"
	redisLib "github.com/dhuki/go-date/pkg/redis"
	"github.com/dhuki/go-date/pkg/validation"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func NewHttpHandler(app *echo.Echo) {
	app.Use(echoMiddleware.TimeoutWithConfig(echoMiddleware.TimeoutConfig{
		Timeout:      config.Conf.App.Timeout,
		ErrorMessage: "timeout",
	}))
	app.Use(echoMiddleware.Recover())
	app.Use(middleware.Logger)

	v1Group := app.Group("/api/v1")

	// dependency
	validationSvc := validation.NewValidation()
	redisLibs := redisLib.NewRedisLibs(redisClient.RedisClient)

	// repository
	repo := repository.NewRepository(database.PostgresDb.Master, database.PostgresDb.Slave)

	// user service
	userSvc := userService.NewUserService(repo, validationSvc, redisLibs)

	// candidate service
	candidateSvc := candidateService.NewCandidateService(repo, redisLibs)

	v1.NewDateHandler(
		userSvc, repo, candidateSvc,
		validationSvc, redisLibs).RegistRoute(v1Group)
}
