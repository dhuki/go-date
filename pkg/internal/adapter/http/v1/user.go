package v1

import (
	"net/http"

	"github.com/dhuki/go-date/pkg/internal/adapter/http/helper"
	"github.com/dhuki/go-date/pkg/internal/adapter/http/v1/model"
	"github.com/labstack/echo/v4"
)

func (handler DateHandler) ListUserRoute(app *echo.Group) {
	app.POST("/user/sign-up", handler.signUp())
	app.POST("/user/login", handler.login())
}

// uploadDocument godoc
// @Summary 	Upload document form-data
// @Description	Upload document form-data
// @Tags		http
// @Accept 		multipart/form-data
// @Param 		file 		formData 	file 	true	"file"
// @Param 		name 		formData 	string 	true	"name"
// @Param 		type 		formData 	string 	true	"type"
// @Param 		path 		formData 	string 	true	"path"
// @Param 		callbackUrl formData	string	false	"callback url"
// @Param 		metadata 	formData 	string 	true 	"metadata of json map<string, any>"
// @Produce		json
// @Success 200 {object} helper.Response{data=DoUploadDocumentResponse,list=interface{}} "Response indicates that the request succeeded and the resources has been fetched and transmitted in the message body"
// @Router /documents/upload-form-data [post]
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
