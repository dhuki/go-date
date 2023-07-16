package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dhuki/go-date/config"
	"github.com/dhuki/go-date/pkg/internal/adapter/http/helper"
	"github.com/dhuki/go-date/pkg/internal/adapter/repository"
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
		}
		logger.Info(ctx, "logger", "%+v", rl)

		return nil
	})
}

func ValidateJWTAccessToken(validationSvc validation.Validation, repo repository.Repository) echo.MiddlewareFunc {
	return echo.MiddlewareFunc(func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			ctx := c.Request().Context()

			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return helper.ResponseError(c, http.StatusUnauthorized, ErrTokenIsEmpty)
			}

			bearerToken := strings.Split(authHeader, " ")
			token := bearerToken[1]
			parsedToken, err := validationSvc.ParseJWTAccessToken(token)
			if err != nil {
				return helper.ResponseError(c, http.StatusUnauthorized, err)
			}

			mapClaim := parsedToken.Claims.(jwt.MapClaims)
			userIDStr := fmt.Sprint(mapClaim["jti"])
			userID, err := strconv.ParseUint(userIDStr, 10, 64)
			if err != nil {
				return helper.ResponseError(c, http.StatusInternalServerError, err)
			}

			user, err := repo.GetUserByID(ctx, userID)
			if err != nil {
				return helper.ResponseError(c, http.StatusInternalServerError, err)
			}

			ctx = context.WithValue(ctx, config.ValueUserIDctx, userIDStr)
			ctx = context.WithValue(ctx, config.ValueUserIDIsPremiumctx, user.IsPremium)
			c.SetRequest(c.Request().WithContext(ctx))
			if err = next(c); err != nil {
				return helper.ResponseError(c, http.StatusInternalServerError, err)
			}
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
				return helper.ResponseError(c, http.StatusInternalServerError, err)
			}
			defer redisLib.Delete(swipeLockingKey)

			isPremiumStr := fmt.Sprint(ctx.Value(config.ValueUserIDIsPremiumctx))
			isPremium, _ := strconv.ParseBool(isPremiumStr)
			if isPremium {
				if err := next(c); err != nil {
					return err
				}
				return nil
			}

			swipeKey := fmt.Sprintf("%s.%s", KeyCountSwipeAction, ctx.Value(config.ValueUserIDctx))
			if value := redisLib.Get(swipeKey); len(value) > 0 {
				currentCount, err := strconv.Atoi(value)
				if err != nil {
					return helper.ResponseError(c, http.StatusInternalServerError, err)
				}
				if currentCount >= config.Conf.RateLimiter.MaxSwipeAction {
					return helper.ResponseError(c, http.StatusTooManyRequests, ErrRateLimiteReachedMaxAttempt)
				}
				if err = next(c); err != nil {
					return err
				}
				redisLib.SetIncr(swipeKey)
				return nil
			}
			if err := next(c); err != nil {
				return err
			}
			if err := redisLib.Set(swipeKey, 1, config.Conf.Redis.CountSwipeActionTTL); err != nil {
				return helper.ResponseError(c, http.StatusInternalServerError, err)
			}
			return nil
		})
	})
}
