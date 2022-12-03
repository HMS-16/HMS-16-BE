package middleware

import (
	"HMS-16-BE/config"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"time"
)

func CreateToken(username, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["role"] = role
	claims["username"] = username
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(12 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Cfg.JWT_SECRET_KEY))
}

func GetJWT(c echo.Context) string {
	token := c.Get("jwt-token").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	role := claims["role"].(string)
	return role
}
