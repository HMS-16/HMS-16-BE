package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func AuthorizationAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role := GetJWT(c)
		if role != "admin" {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "url access for admin",
			})
		}
		return next(c)
	}
}
