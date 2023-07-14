package router

import (
	"context"
	"net/http"

	httpHandler "github.com/dhuki/go-date/pkg/internal/adapter/http"
	"github.com/labstack/echo/v4"
)

type EchoRouter interface {
	Start(ctx context.Context) (err error)
	Stop(ctx context.Context) (err error)
}

type EchoRouterImpl struct {
	echoRouter *echo.Echo
	addr       string
}

func NewEchoRouter(echoRouter *echo.Echo, addr string) EchoRouter {
	return EchoRouterImpl{
		echoRouter: echoRouter,
		addr:       addr,
	}
}

func (e EchoRouterImpl) Start(ctx context.Context) (err error) {
	httpHandler.NewHttpHandler(e.echoRouter)
	if err = e.echoRouter.Start(e.addr); err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (e EchoRouterImpl) Stop(ctx context.Context) (err error) {
	return e.echoRouter.Shutdown(ctx)
}
