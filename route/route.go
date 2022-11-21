package route

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Init(e *echo.Echo) {
	v1 := e.Group("/v1")

	v1.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "success",
		})
	})
}
