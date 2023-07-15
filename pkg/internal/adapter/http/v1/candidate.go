package v1

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/dhuki/go-date/pkg/internal/adapter/http/helper"
	"github.com/dhuki/go-date/pkg/middleware"
	"github.com/labstack/echo/v4"
)

func (handler DateHandler) ListCandidateRoute(app *echo.Group) {
	v1GroupCandidate := app.Group("/candidate")
	v1GroupCandidate.Use(middleware.ValidateJWTAccessToken(handler.validationService))

	v1GroupCandidate.GET("/candidate", handler.getListCandidate())
	v1GroupCandidate.POST("/candidate/swipe/:candidateId", middleware.RateLimiter(handler.redisLibs)(handler.swipeAction()))
}

func (d DateHandler) swipeAction() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		candidateIDStr := c.Param("candidateId")
		if candidateIDStr == "" {
			return helper.ResponseError(c, http.StatusBadRequest, errors.New("todo fixing"))
		}

		swipeDirection := c.QueryParam("to")
		if swipeDirection == "" {
			return helper.ResponseError(c, http.StatusBadRequest, errors.New("todo fixing"))
		}

		candidateID, err := strconv.ParseUint(candidateIDStr, 10, 64)
		if err != nil {
			return helper.ResponseError(c, http.StatusBadRequest, err)
		}

		if err = d.candidateService.SwipeAction(ctx, candidateID, swipeDirection); err != nil {
			return helper.ResponseError(c, http.StatusInternalServerError, err)
		}

		return helper.ResponseSuccess(c, "success", nil)
	}
}

func (d DateHandler) getListCandidate() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		limitStr := c.QueryParam("limit")
		if limitStr == "" {
			return helper.ResponseError(c, http.StatusBadRequest, errors.New("todo fixing"))
		}

		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			return helper.ResponseError(c, http.StatusBadRequest, errors.New("todo fixing"))
		}

		data, err := d.candidateService.GetListCandidate(ctx, limit)
		if err != nil {
			return helper.ResponseError(c, http.StatusInternalServerError, err)
		}

		return helper.ResponseSuccessPagination(c, "success", data.CandidateList, data.Page)
	}
}
