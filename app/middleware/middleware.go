package middleware

import (
	"net/http"

	"github.com/labstack/echo"
)

// AppMiddleware represent the app middleware name
type AppMiddleware struct {
	AppName string
}

// CORS middleware config
func (am *AppMiddleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Server", am.AppName)
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST")
		c.Response().Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		c.Response().Header().Set("Content-Type", "application/json")

		if c.Request().Method == "OPTIONS" {
			return c.String(http.StatusOK, "")
		}

		return next(c)
	}
}

// InitAppMiddleware initialization
func InitAppMiddleware(appName string) *AppMiddleware {
	return &AppMiddleware{
		AppName: appName,
	}
}
