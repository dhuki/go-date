package v1

import "github.com/labstack/echo/v4"

func (handler DateHandler) ListCandidateRoute(app *echo.Group) {
	app.POST("/candidate/swipe", handler.swipeAction())
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
func (d DateHandler) swipeAction() echo.HandlerFunc {
	return func(c echo.Context) error {
		d.candidateService.SwipeAction(c.Request().Context())
		return nil
	}
}
