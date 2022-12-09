package route

import (
	"HMS-16-BE/config"
	adminctrl "HMS-16-BE/controller/admin"
	patientctrl "HMS-16-BE/controller/patient"
	profilectrl "HMS-16-BE/controller/profile"
	schedulectrl "HMS-16-BE/controller/schedule"
	userctrl "HMS-16-BE/controller/user"
	adminrepo "HMS-16-BE/repository/admin"
	patientrepo "HMS-16-BE/repository/patient"
	profilerepo "HMS-16-BE/repository/profile"
	schedulerepo "HMS-16-BE/repository/schedule"
	userrepo "HMS-16-BE/repository/user"
	adminuc "HMS-16-BE/usecase/admin"
	patientuc "HMS-16-BE/usecase/patient"
	profileuc "HMS-16-BE/usecase/profile"
	scheduleuc "HMS-16-BE/usecase/schedule"
	useruc "HMS-16-BE/usecase/user"
	"HMS-16-BE/util/middleware"
	"database/sql"
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
	"os"
)

func Init(e *echo.Echo, db *sql.DB) {
	adminRepo := adminrepo.NewAdminRepository(db)
	userRepo := userrepo.NewUserRepository(db)
	patientRepo := patientrepo.NewPatientRepository(db)
	guardianRepo := patientrepo.NewGuardianRepository(db)
	doctorRepo := profilerepo.NewDoctorRepository(db)
	nurseRepo := profilerepo.NewNurseRepository(db)
	shiftRepo := schedulerepo.NewShiftRepository(db)

	adminUC := adminuc.NewAdminUsecase(adminRepo)
	userUC := useruc.NewUserUsecase(userRepo)
	patientUC := patientuc.NewPatientUsecase(patientRepo, guardianRepo)
	guardianUC := patientuc.NewGuardianUSecase(guardianRepo)
	doctorUC := profileuc.NewDoctorUsecase(doctorRepo)
	nurseUC := profileuc.NewNurseUsecase(nurseRepo)
	shiftUC := scheduleuc.NewShiftUsecase(shiftRepo)

	adminCtrl := adminctrl.NewAdminController(adminUC)
	userCtrl := userctrl.NewUserController(userUC)
	patientCtrl := patientctrl.NewPatientController(patientUC)
	guardianCtrl := patientctrl.NewGuardianController(guardianUC)
	doctorCtrl := profilectrl.NewDoctorController(doctorUC)
	nurseCtrl := profilectrl.NewNurseController(nurseUC)
	shiftCtrl := schedulectrl.NewShiftController(shiftUC)

	secretJWT := os.Getenv("JWT_SECRET_KEY")
	if secretJWT == "" {
		secretJWT = config.Cfg.JWT_SECRET_KEY
	}

	middleware.LogMiddleware(e)
	v1 := e.Group("/v1")
	adminV1 := v1.Group("/admins")
	adminV1.POST("/signup", adminCtrl.Create)
	adminV1.POST("/login", adminCtrl.Login)

	adminV1JWT := e.Group("/v1/admins")
	adminV1JWT.Use(mid.JWTWithConfig(mid.JWTConfig{
		SigningKey: []byte(secretJWT),
		ContextKey: "jwt-token",
	}))
	adminV1JWT.Use(middleware.AuthorizationAdmin)
	adminV1JWT.GET("/:id", adminCtrl.GetById)
	adminV1JWT.PUT("/:id", adminCtrl.Update)
	adminV1JWT.DELETE("/:id", adminCtrl.Delete)

	userV1 := e.Group("/v1")
	userV1.POST("/login", userCtrl.Login)
	userV1JWT := e.Group("/v1")
	userV1JWT.Use(mid.JWTWithConfig(mid.JWTConfig{
		SigningKey: []byte(secretJWT),
		ContextKey: "jwt-token",
	}))
	userV1JWTAdmin := userV1JWT.Group("")
	userV1JWTAdmin.Use(middleware.AuthorizationAdmin)
	userV1JWTAdmin.POST("/register", userCtrl.Create)
	userV1JWTAdmin.GET("/accounts", userCtrl.GetAll)
	userV1JWT.GET("/accounts/:id", userCtrl.GetById)
	userV1JWT.PUT("/accounts/:id", userCtrl.Update)
	userV1JWT.DELETE("/accounts/:id", userCtrl.Delete)

	patients := e.Group("/v1/patients")
	patients.Use(mid.JWTWithConfig(mid.JWTConfig{
		SigningKey: []byte(secretJWT),
		ContextKey: "jwt-token",
	}))
	patients.Use(middleware.AuthorizationAdmin)
	patients.GET("", patientCtrl.GetAll)
	patients.GET("/:id", patientCtrl.GetById)
	patients.POST("", patientCtrl.Create)
	patients.PUT("/:id", patientCtrl.Update)
	patients.DELETE("/:id", patientCtrl.Delete)

	guardian := patients.Group("/guardians")
	guardian.GET("/:id", guardianCtrl.GetById)   //id = guardian id
	guardian.POST("/:id", guardianCtrl.Create)   //id = patient id
	guardian.PUT("/:id", guardianCtrl.Update)    //guardian id
	guardian.DELETE("/:id", guardianCtrl.Delete) //guardian id

	doctor := v1.Group("/doctors")
	doctor.Use(mid.JWTWithConfig(mid.JWTConfig{
		SigningKey: []byte(secretJWT),
		ContextKey: "jwt-token",
	}))
	doctor.Use(middleware.AuthorizationDoctor)
	doctor.GET("/all", doctorCtrl.GetAll)
	doctor.GET("", doctorCtrl.GetById)
	doctor.POST("", doctorCtrl.Create)
	doctor.PUT("", doctorCtrl.Update)
	doctor.DELETE("", doctorCtrl.Delete)

	nurse := v1.Group("/nurses")
	nurse.Use(mid.JWTWithConfig(mid.JWTConfig{
		SigningKey: []byte(secretJWT),
		ContextKey: "jwt-token",
	}))
	nurse.Use(middleware.AuthorizationNurse)
	nurse.GET("/all", nurseCtrl.GetAll)
	nurse.GET("", nurseCtrl.GetById)
	nurse.POST("", nurseCtrl.Create)
	nurse.PUT("", nurseCtrl.Update)
	nurse.DELETE("", nurseCtrl.Delete)

	shiftDoctor := doctor.Group("")
	shiftDoctor.GET("", shiftCtrl.GetAllByUserId)
	shiftDoctor.GET("/:id", shiftCtrl.GetById)
	shiftDoctor.POST("", shiftCtrl.Create)
	shiftDoctor.POST("/:id", shiftCtrl.Update)
	shiftDoctor.DELETE("/:id", shiftCtrl.Delete)

	shiftNurse := nurse.Group("")
	shiftNurse.GET("", shiftCtrl.GetAllByUserId)
	shiftNurse.GET("/:id", shiftCtrl.GetById)
	shiftNurse.POST("", shiftCtrl.Create)
	shiftNurse.POST("/:id", shiftCtrl.Update)
	shiftNurse.DELETE("/:id", shiftCtrl.Delete)
}
