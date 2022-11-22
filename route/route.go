package route

import (
	controller "HMS-16-BE/controller/admin"
	"HMS-16-BE/repository/admin"
	"HMS-16-BE/usecase/admin"
	"database/sql"
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, db *sql.DB) {
	adminRepo := repository.NewAdminRepository(db)
	adminUC := usecase.NewAdminUsecase(adminRepo)
	adminCtrl := controller.NewAdminController(adminUC)

	v1 := e.Group("/v1")
	adminV1 := v1.Group("/admins")
	adminV1.POST("/signup", adminCtrl.Create)
	adminV1.POST("/login", adminCtrl.Login)
	adminV1.GET("/:id", adminCtrl.GetById)
	adminV1.PUT("/:id", adminCtrl.Update)
	adminV1.DELETE("/:id", adminCtrl.Delete)
}
