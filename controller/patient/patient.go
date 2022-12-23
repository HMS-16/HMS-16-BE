package patient

import (
	"HMS-16-BE/model"
	"HMS-16-BE/usecase/patient"
	"HMS-16-BE/util/middleware"
	"HMS-16-BE/util/uuid"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type patientController struct {
	patient patient.PatientUsecase
}

func NewPatientController(p patient.PatientUsecase) *patientController {
	return &patientController{p}
}

func (p *patientController) GetAll(c echo.Context) error {
	patients, err := p.patient.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    patients,
	})
}

func (p *patientController) GetAllCards(c echo.Context) error {
	patients, err := p.patient.GetAllCards()
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    patients,
	})
}

func (p *patientController) GetById(c echo.Context) error {
	id := c.Param("id")

	patient, err := p.patient.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    patient,
	})
}

func (p *patientController) Create(c echo.Context) error {
	var patient model.Patients
	c.Bind(&patient)
	patient.Id = uuid.CreateUUID()
	patient.CreatedAt = time.Now()
	patient.UpdatedAt = patient.CreatedAt
	patient.AdminId = middleware.GetIdJWT(c)

	validate := validator.New()
	err := validate.Struct(&patient)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	err = p.patient.Create(patient)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}

func (p *patientController) Update(c echo.Context) error {
	var patient model.Patients
	id := c.Param("id")
	c.Bind(&patient)
	patient.Id = id
	patient.UpdatedAt = time.Now()

	validate := validator.New()
	err := validate.Struct(&patient)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	err = p.patient.Update(patient)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}

func (p *patientController) UpdateEndCase(c echo.Context) error {
	id := c.Param("id")

	err := p.patient.UpdateEndCase(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}

func (p *patientController) Delete(c echo.Context) error {
	id := c.Param("id")

	err := p.patient.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}
