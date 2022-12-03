package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LogMiddleware(c *echo.Echo) {
	c.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}" +
			"time_unix=${time_unix}, user_agent=${user_agent}" + "\n",
	}))
}
