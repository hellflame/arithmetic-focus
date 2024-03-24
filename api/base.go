package api

import (
	"embed"

	"github.com/hellflame/arithmetic-focus/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

var router *echo.Echo

type Response struct {
	Code    int
	Message string
	Content any
}

func init() {
	router = echo.New()

	router.Use(middleware.CORS())
	router.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogMethod: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			utils.Logger.Info("request",
				zap.String("method", v.Method),
				zap.String("uri", v.URI),
				zap.Int("status", v.Status),
			)
			return nil
		},
	}))
}

func RegisterAPI(pages *embed.FS) *echo.Echo {
	bindPages(pages)
	bindExpression()
	bindRecord()
	return router
}

func responseSimpleMessage(c echo.Context, code int, message string) error {
	return c.JSONBlob(200, utils.LowerCamelMarshal(Response{Code: code, Message: message, Content: ""}))
}

func responseLowerJson(c echo.Context, entity any) error {
	return c.JSONBlob(200, utils.LowerCamelMarshal(Response{Content: entity}))
}
