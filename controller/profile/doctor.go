package profile

import (
	"HMS-16-BE/model"
	"HMS-16-BE/usecase/profile"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type doctorController struct {
	doctor profile.DoctorUsecase
}

func NewDoctorController(d profile.DoctorUsecase) *doctorController {
	return &doctorController{d}
}

func (d *doctorController) GetAll(c echo.Context) error {
	doctors, err := d.doctor.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    doctors,
	})
}

func (d *doctorController) GetAllCards(c echo.Context) error {
	doctors, err := d.doctor.GetAllCards()
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    doctors,
	})
}

func (d *doctorController) GetById(c echo.Context) error {
	id := c.Param("id")

	doctor, err := d.doctor.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
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

	validate := validator.New()
	err := validate.Struct(&doctor)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	err = d.doctor.Create(doctor)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}

func (d *doctorController) Update(c echo.Context) error {
	var doctor model.Doctors
	c.Bind(&doctor)
	doctor.StrNum = c.Param("id")
	doctor.UpdatedAt = time.Now()

	validate := validator.New()
	err := validate.Struct(doctor)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	err = d.doctor.Update(doctor)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}

func (d *doctorController) Delete(c echo.Context) error {
	id := c.Param("id")

	err := d.doctor.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}
