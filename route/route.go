package route

import (
	"HMS-16-BE/config"
	adminctrl "HMS-16-BE/controller/admin"
	"HMS-16-BE/controller/outpatientSession"
	patientctrl "HMS-16-BE/controller/patient"
	profilectrl "HMS-16-BE/controller/profile"
	shiftctrl "HMS-16-BE/controller/shift"
	userctrl "HMS-16-BE/controller/user"
	adminrepo "HMS-16-BE/repository/admin"
	outpatientSessionrepo "HMS-16-BE/repository/outpatientSession"
	patientrepo "HMS-16-BE/repository/patient"
	profilerepo "HMS-16-BE/repository/profile"
	shiftrepo "HMS-16-BE/repository/shift"
	userrepo "HMS-16-BE/repository/user"
	adminuc "HMS-16-BE/usecase/admin"
	outpatientSessionuc "HMS-16-BE/usecase/outpatientSession"
	patientuc "HMS-16-BE/usecase/patient"
	profileuc "HMS-16-BE/usecase/profile"
	shiftuc "HMS-16-BE/usecase/shift"
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
	doctorRepo := profilerepo.NewDoctorRepository(db)
	nurseRepo := profilerepo.NewNurseRepository(db)
	shiftRepo := shiftrepo.NewShiftRepository(db)
	timeRepo := shiftrepo.NewTimeRepository(db)
	dayRepo := shiftrepo.NewDayRepository(db)
	scheduleRepo := outpatientSessionrepo.NewScheduleRepository(db)
	conditionRepo := outpatientSessionrepo.NewConditionRepository(db)
	diagnoseRepo := outpatientSessionrepo.NewDiagnoseRepository(db)

	adminUC := adminuc.NewAdminUsecase(adminRepo)
	userUC := useruc.NewUserUsecase(userRepo, doctorRepo, nurseRepo)
	patientUC := patientuc.NewPatientUsecase(patientRepo)
	doctorUC := profileuc.NewDoctorUsecase(doctorRepo)
	nurseUC := profileuc.NewNurseUsecase(nurseRepo)
	shiftUC := shiftuc.NewShiftUsecase(shiftRepo, dayRepo, timeRepo)
	timeUC := shiftuc.NewtimeUsecase(timeRepo)
	dayUC := shiftuc.NewDayUsecase(dayRepo)
	scheduleUC := outpatientSessionuc.NewScheduleUsecase(scheduleRepo, conditionRepo, diagnoseRepo, patientRepo,
		userRepo, shiftRepo)
	conditionUC := outpatientSessionuc.NewConditionUsecase(conditionRepo, scheduleRepo, userRepo)
	diagnoseUC := outpatientSessionuc.NewDiagnoseUseCase(diagnoseRepo, scheduleRepo, userRepo)

	adminCtrl := adminctrl.NewAdminController(adminUC)
	userCtrl := userctrl.NewUserController(userUC)
	patientCtrl := patientctrl.NewPatientController(patientUC)
	doctorCtrl := profilectrl.NewDoctorController(doctorUC)
	nurseCtrl := profilectrl.NewNurseController(nurseUC)
	shiftCtrl := shiftctrl.NewShiftController(shiftUC)
	timeCtrl := shiftctrl.NewTimeController(timeUC)
	dayCtrl := shiftctrl.NewDayController(dayUC)
	scheduleCtrl := outpatientSession.NewScheduleController(scheduleUC)
	conditionCtrl := outpatientSession.NewConditionController(conditionUC)
	diagnoseCtrl := outpatientSession.NewDiagnoseController(diagnoseUC)

	secretJWT := os.Getenv("JWT_SECRET_KEY")
	if secretJWT == "" {
		secretJWT = config.Cfg.JWT_SECRET_KEY
	}

	middleware.LogMiddleware(e)
	e.Use(mid.CORSWithConfig(mid.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization,
			echo.HeaderAcceptEncoding, echo.HeaderXCSRFToken, echo.HeaderContentLength},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
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
	userV1JWT.GET("/accounts", userCtrl.GetAll)
	userV1JWT.GET("/accounts/:id", userCtrl.GetById)
	userV1JWT.PUT("/accounts/:id", userCtrl.Update)
	userV1JWT.PUT("/accounts/password/:id", userCtrl.UpdatePassword)
	userV1JWT.DELETE("/accounts/:id", userCtrl.Delete)

	patients := e.Group("/v1/patients")
	patients.Use(mid.JWTWithConfig(mid.JWTConfig{
		SigningKey: []byte(secretJWT),
		ContextKey: "jwt-token",
	}))
	patientsAdmin := patients.Group("")
	patientsAdmin.Use(middleware.AuthorizationAdmin)
	patientsAdmin.POST("", patientCtrl.Create)
	patientsAdmin.PUT("/:id", patientCtrl.Update)
	patientsAdmin.DELETE("/:id", patientCtrl.Delete)
	patients.GET("", patientCtrl.GetAll)
	patients.GET("/cards", patientCtrl.GetAllCards)
	patients.GET("/:id", patientCtrl.GetById)
	patients.PUT("/endcase/:id", patientCtrl.UpdateEndCase) //id patient

	doctor := v1.Group("/doctors")
	doctor.Use(mid.JWTWithConfig(mid.JWTConfig{
		SigningKey: []byte(secretJWT),
		ContextKey: "jwt-token",
	}))
	doctor.Use(middleware.AuthorizationDoctor)
	doctor.GET("/all", doctorCtrl.GetAll)
	doctor.GET("/all/cards", doctorCtrl.GetAllCards)
	doctor.GET("/:id", doctorCtrl.GetById)
	doctor.POST("", doctorCtrl.Create)
	doctor.PUT("/:id", doctorCtrl.Update)
	doctor.DELETE("/:id", doctorCtrl.Delete)

	nurse := v1.Group("/nurses")
	nurse.Use(mid.JWTWithConfig(mid.JWTConfig{
		SigningKey: []byte(secretJWT),
		ContextKey: "jwt-token",
	}))
	nurse.Use(middleware.AuthorizationNurse)
	nurse.GET("/all", nurseCtrl.GetAll)
	nurse.GET("/all/cards", nurseCtrl.GetAllCards)
	nurse.GET("/:id", nurseCtrl.GetById)
	nurse.POST("", nurseCtrl.Create)
	nurse.PUT("/:id", nurseCtrl.Update)
	nurse.DELETE("/:id", nurseCtrl.Delete)

	shiftDoctor := doctor.Group("/shifts")
	shiftDoctor.GET("", shiftCtrl.GetAll)
	shiftDoctor.GET("/all/:id", shiftCtrl.GetAllByUserId)
	shiftDoctor.GET("/:id", shiftCtrl.GetById)
	shiftDoctor.POST("", shiftCtrl.Create)
	shiftDoctor.PUT("/:id", shiftCtrl.Update)
	shiftDoctor.DELETE("/:id", shiftCtrl.Delete)

	shiftNurse := nurse.Group("/shifts")
	shiftNurse.GET("", shiftCtrl.GetAll)
	shiftNurse.GET("/all/:id", shiftCtrl.GetAllByUserId)
	shiftNurse.GET("/:id", shiftCtrl.GetById)
	shiftNurse.POST("", shiftCtrl.Create)
	shiftNurse.PUT("/:id", shiftCtrl.Update)
	shiftNurse.DELETE("/:id", shiftCtrl.Delete)

	time := v1.Group("/shifts/times")
	time.GET("", timeCtrl.GetAll)
	time.GET("/:id", timeCtrl.GetById)
	time.POST("", timeCtrl.Create)
	time.PUT("/:id", timeCtrl.Update)
	time.DELETE("/:id", timeCtrl.Delete)

	day := v1.Group("/shifts/day")
	day.GET("", dayCtrl.GetAll)
	day.GET("/:id", dayCtrl.GetById)
	day.POST("", dayCtrl.Create)
	day.PUT("/:id", dayCtrl.Update)
	day.DELETE("/:id", dayCtrl.Delete)

	appointment := e.Group("/v1/appointment")
	appointment.Use(mid.JWTWithConfig(mid.JWTConfig{
		SigningKey: []byte(secretJWT),
		ContextKey: "jwt-token",
	}))
	appointmentAdmin := appointment.Group("")
	appointmentAdmin.Use(middleware.AuthorizationAdmin)
	appointmentAdmin.POST("", scheduleCtrl.Create)
	appointment.GET("", scheduleCtrl.GetAll)
	appointment.GET("/all", scheduleCtrl.GetAllByDay)
	appointment.GET("/all/cards", scheduleCtrl.GetAllCardByDay)
	appointment.GET("/:id", scheduleCtrl.GetByScheduleId)
	appointment.GET("/patient/:id", scheduleCtrl.GetAllByPatient)
	appointment.GET("/patient/detail/:id", scheduleCtrl.GetDetailByPatient)
	appointment.PUT("/change/doctor/:id", scheduleCtrl.UpdateDoctor)
	appointment.PUT("/change/nurse/:id", scheduleCtrl.UpdateNurse)
	appointment.PUT("/change/status/:id", scheduleCtrl.UpdateStatus)
	appointment.PUT("/change/date/:id", scheduleCtrl.Update)
	appointment.DELETE("/:id", scheduleCtrl.Delete)

	condition := e.Group("/v1/conditions")
	condition.Use(mid.JWTWithConfig(mid.JWTConfig{
		SigningKey: []byte(secretJWT),
		ContextKey: "jwt-token",
	}))
	condition.Use(middleware.AuthorizationNurse)
	condition.POST("/:id", conditionCtrl.Create) //patientid
	condition.GET("/:id", conditionCtrl.GetById)
	condition.GET("/patients/:id", conditionCtrl.GetAllByPatient)

	diagnose := e.Group("/v1/diagnoses")
	diagnose.Use(mid.JWTWithConfig(mid.JWTConfig{
		SigningKey: []byte(secretJWT),
		ContextKey: "jwt-token",
	}))
	diagnose.Use(middleware.AuthorizationDoctor)
	diagnose.POST("/:id", diagnoseCtrl.Create) //patientid
	diagnose.GET("/:id", diagnoseCtrl.GetById)
	diagnose.GET("/patients/:id", diagnoseCtrl.GetAllByPatient)
}
