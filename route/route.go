package route

import (
	"HMS-16-BE/config"
	"HMS-16-BE/controller/admin"
	"HMS-16-BE/repository/admin"
	"HMS-16-BE/usecase/admin"
	"HMS-16-BE/util/middleware"
	"database/sql"
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func Init(e *echo.Echo, db *sql.DB) {
	adminRepo := repository.NewAdminRepository(db)
	adminUC := usecase.NewAdminUsecase(adminRepo)
	adminCtrl := controller.NewAdminController(adminUC)

	middleware.LogMiddleware(e)
	v1 := e.Group("/v1")
	adminV1 := v1.Group("/admins")
	adminV1.POST("/signup", adminCtrl.Create)
	adminV1.POST("/login", adminCtrl.Login)

	adminV1JWT := adminV1
	adminV1JWT.Use(mid.JWT([]byte(config.Cfg.JWT_SECRET_KEY)))
	adminV1JWT.GET("/:id", adminCtrl.GetById)
	adminV1JWT.PUT("/:id", adminCtrl.Update)
	adminV1JWT.DELETE("/:id", adminCtrl.Delete)
}
