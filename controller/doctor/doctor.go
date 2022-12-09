package doctor

import (
	"HMS-16-BE/model"
	"HMS-16-BE/usecase/doctor"
	"HMS-16-BE/util/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type doctorController struct {
	doctor doctor.DoctorUsecase
}

func NewDoctorController(d doctor.DoctorUsecase) *doctorController {
	return &doctorController{d}
}

func (d *doctorController) GetAll(c echo.Context) error {
	doctors, err := d.doctor.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    doctors,
	})
}

func (d *doctorController) GetById(c echo.Context) error {
	id := middleware.GetIdJWT(c)

	doctor, err := d.doctor.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    doctor,
	})
}

func (d *doctorController) Create(c echo.Context) error {
	var doctor model.Doctors
	c.Bind(&doctor)
	doctor.CreatedAt = time.Now()
	doctor.UpdatedAt = doctor.CreatedAt
	doctor.UserId = middleware.GetIdJWT(c)

	validate := validator.New()
	err := validate.Struct(&doctor)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = d.doctor.Create(doctor)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}

func (p *doctorController) Update(c echo.Context) error {
	var doctor model.Doctors
	c.Bind(&doctor)
	doctor.UserId = middleware.GetIdJWT(c)
	doctor.UpdatedAt = time.Now()

	validate := validator.New()
	err := validate.Struct(doctor)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = p.doctor.Update(doctor)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}

func (p *doctorController) Delete(c echo.Context) error {
	id := middleware.GetIdJWT(c)

	err := p.doctor.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}
