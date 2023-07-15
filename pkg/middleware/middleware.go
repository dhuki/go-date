package middleware

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dhuki/go-date/config"
	"github.com/dhuki/go-date/pkg/logger"
	"github.com/dhuki/go-date/pkg/redis"
	"github.com/dhuki/go-date/pkg/validation"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		ctx := c.Request().Context()
		start := time.Now()
		if err := next(c); err != nil {
			c.Error(err)
		}
		rl := requestLog{
			Timestamp:     start,
			CorrelationID: "",
			Method:        c.Request().Method,
			URL:           c.Request().URL.RequestURI(),
			Status:        c.Response().Status,
			ResponseTime:  time.Since(start).Seconds(),
			ResponseSize:  c.Response().Size,
			// ReqBody:       ctx.Value(utils.LOG_REQ_BODY),
		}

		logger.Info(ctx, "logger", "%+v", rl)

		// b, err := json.Marshal(&rl)
		// if err == nil {
		// 	logger.Info(ctx, "logger", "%+v", rl)
		// }

		return nil
	})
}

func ValidateJWTAccessToken(validationSvc validation.Validation) echo.MiddlewareFunc {
	return echo.MiddlewareFunc(func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			ctx := c.Request().Context()

			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return errors.New("error")
			}

			bearerToken := strings.Split(authHeader, " ")
			token := bearerToken[1]

			parsedToken, err := validationSvc.ParseJWTAccessToken(token)
			if err != nil {
				return err
			}

			mapClaim := parsedToken.Claims.(jwt.MapClaims)
			ctx = context.WithValue(ctx, config.ValueUserIDctx, mapClaim["jti"])
			c.SetRequest(c.Request().WithContext(ctx))
			next(c)
			return nil
		})
	})
}

func RateLimiter(redisLib redis.Redis) echo.MiddlewareFunc {
	return echo.MiddlewareFunc(func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			ctx := c.Request().Context()

			swipeLockingKey := fmt.Sprintf("%s.%s", KeyLockingSwipeAction, ctx.Value(config.ValueUserIDctx))
			if err := redisLib.SetLockingKey(swipeLockingKey, true, config.Conf.Redis.LockingSwipeActionTTL); err != nil {
				return err
			}
			defer redisLib.Delete(swipeLockingKey)

			swipeKey := fmt.Sprintf("%s.%s", KeyCountSwipeAction, ctx.Value(config.ValueUserIDctx))
			if value := redisLib.Get(swipeKey); len(value) > 0 {
				currentCount, err := strconv.Atoi(value)
				if err != nil {
					return err
				}
				if currentCount >= config.Conf.RateLimiter.MaxSwipeAction {
					return ErrRateLimiteReachedMaxAttempt
				}
				if err = next(c); err != nil {
					return err
				}
				redisLib.SetIncr(swipeKey)
				return nil
			}

			if err := redisLib.Set(swipeKey, 1, config.Conf.Redis.CountSwipeActionTTL); err != nil {
				return err
			}
			return nil
		})
	})
}
