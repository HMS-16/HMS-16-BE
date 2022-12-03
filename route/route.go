package route

import (
	"HMS-16-BE/config"
	adminctrl "HMS-16-BE/controller/admin"
	userctrl "HMS-16-BE/controller/user"
	adminrepo "HMS-16-BE/repository/admin"
	userrepo "HMS-16-BE/repository/user"
	adminuc "HMS-16-BE/usecase/admin"
	useruc "HMS-16-BE/usecase/user"
	"HMS-16-BE/util/middleware"
	"database/sql"
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func Init(e *echo.Echo, db *sql.DB) {
	adminRepo := adminrepo.NewAdminRepository(db)
	userRepo := userrepo.NewUserRepository(db)

	adminUC := adminuc.NewAdminUsecase(adminRepo)
	userUC := useruc.NewUserUsecase(userRepo)

	adminCtrl := adminctrl.NewAdminController(adminUC)
	userCtrl := userctrl.NewUserController(userUC)

	middleware.LogMiddleware(e)
	v1 := e.Group("/v1")
	adminV1 := v1.Group("/admins")
	adminV1.POST("/signup", adminCtrl.Create)
	adminV1.POST("/login", adminCtrl.Login)

	adminV1JWT := adminV1
	adminV1JWT.Use(mid.JWTWithConfig(mid.JWTConfig{
		SigningKey: []byte(config.Cfg.JWT_SECRET_KEY),
		ContextKey: "jwt-token",
	}))
	adminV1JWT.Use(middleware.AuthorizationAdmin)
	adminV1JWT.GET("/:id", adminCtrl.GetById)
	adminV1JWT.PUT("/:id", adminCtrl.Update)
	adminV1JWT.DELETE("/:id", adminCtrl.Delete)

	userV1 := v1
	userV1.POST("/login", userCtrl.Login)
	userV1JWT := userV1
	userV1JWT.Use(mid.JWTWithConfig(mid.JWTConfig{
		SigningKey: []byte(config.Cfg.JWT_SECRET_KEY),
		ContextKey: "jwt-token",
	}))
	userV1JWTAdmin := userV1JWT
	userV1JWTAdmin.Use(middleware.AuthorizationAdmin)
	userV1JWTAdmin.POST("/register", userCtrl.Create)
	userV1JWTAdmin.GET("/accounts", userCtrl.GetAll)
	userV1JWT.GET("/accounts/:id", userCtrl.GetById)
	userV1JWT.PUT("/accounts/:id", userCtrl.Update)
	userV1JWT.DELETE("/accounts/:id", userCtrl.Delete)
}
