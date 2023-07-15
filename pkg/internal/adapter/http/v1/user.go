package v1

import (
	"net/http"

	"github.com/dhuki/go-date/pkg/internal/adapter/http/helper"
	"github.com/dhuki/go-date/pkg/internal/adapter/http/v1/model"
	"github.com/labstack/echo/v4"
)

func (handler DateHandler) ListUserRoute(app *echo.Group) {
	v1GroupUser := app.Group("/user")

	v1GroupUser.POST("/sign-up", handler.signUp())
	v1GroupUser.POST("/login", handler.login())
}

func (d DateHandler) signUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var reqBody model.CreateUserRequest
		if err := c.Bind(&reqBody); err != nil {
			return helper.ResponseError(c, http.StatusBadRequest, err)
		}

		if err := d.userService.SignUp(ctx, reqBody); err != nil {
			return helper.ResponseError(c, http.StatusInternalServerError, err)
		}

		return helper.ResponseSuccess(c, "success", nil)
	}
}

func (d DateHandler) login() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var reqBody model.LoginRequest
		if err := c.Bind(&reqBody); err != nil {
			return helper.ResponseError(c, http.StatusBadRequest, err)
		}

		resp, err := d.userService.Login(ctx, reqBody)
		if err != nil {
			return helper.ResponseError(c, http.StatusInternalServerError, err)
		}

		return helper.ResponseSuccess(c, "success", resp)
	}
}
