package middleware

import (
	"HMS-16-BE/config"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"os"
	"time"
)

func CreateToken(id, username, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["role"] = role
	claims["username"] = username
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(12 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		secretKey = config.Cfg.JWT_SECRET_KEY
	}
	return token.SignedString([]byte(secretKey))
}

func GetRoleJWT(c echo.Context) string {
	token := c.Get("jwt-token").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	role := claims["role"].(string)
	return role
}

func GetIdJWT(c echo.Context) string {
	token := c.Get("jwt-token").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	return id
}
