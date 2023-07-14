package middleware

import (
	"time"

	"github.com/dhuki/go-date/pkg/logger"
	"github.com/labstack/echo/v4"
)

func Logger() echo.MiddlewareFunc {
	return echo.MiddlewareFunc(func(next echo.HandlerFunc) echo.HandlerFunc {
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
	})
}
