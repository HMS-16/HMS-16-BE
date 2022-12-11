package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func AuthorizationAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role := GetRoleJWT(c)
		if role != "admin" {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "url access for admin",
			})
		}
		return next(c)
	}
}

func AuthorizationDoctor(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role := GetRoleJWT(c)
		if role != "doctor" && role != "admin" {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "url access for doctor",
			})
		}
		return next(c)
	}
}

func AuthorizationNurse(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role := GetRoleJWT(c)
		if role != "nurse" && role != "admin" {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "url access for nurse",
			})
		}
		return next(c)
	}
}
