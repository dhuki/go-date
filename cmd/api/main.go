package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/dhuki/go-date/config"
	"github.com/dhuki/go-date/config/consul"
	"github.com/dhuki/go-date/config/database"
	"github.com/dhuki/go-date/config/redis"
	"github.com/dhuki/go-date/pkg/logger"
	"github.com/dhuki/go-date/pkg/router"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	env := os.Getenv("ENV")
	if env == "" {
		env = "LOCAL"
	}
	flag.StringVar(&config.Conf.App.Env, "config_name", env, "define environment")
	flag.Parse()

	// init config
	consul.InitConsul(ctx, config.Conf.App.Env)

	// init postgres database
	if err := database.InitPostgres(&config.Conf.ConnDatabase); err != nil {
		logger.Fatal(ctx, "database.InitPostgres", "Error connect to database, err : %v", err)
	}

	// init redis
	if err := redis.InitRedis(&config.Conf.Redis); err != nil {
		logger.Fatal(ctx, "redis.InitRedis", "Error connect to redis, err : %v", err)
	}

	// init router
	echoRoute := router.NewEchoRouter(echo.New(), fmt.Sprintf(":%d", config.Conf.App.Port))
	idleConsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		logrus.Infof("We received an interrupt signal, shutting down service")
		if err := echoRoute.Stop(ctx); err != nil {
			logger.Fatal(ctx, "route stop", "Error stopping service go-date, err : %v", err)
		}
		logger.Info(ctx, "route stop", "Success stopping service go-date")
		close(idleConsClosed)
	}()
	logger.Info(ctx, "route start", "Success start service go-date listening on port :%d", config.Conf.App.Port)
	if err := echoRoute.Start(ctx); err != nil {
		logger.Fatal(ctx, "route start", "Error starting service go-date, err : %v", err)
	}
	<-idleConsClosed
}
