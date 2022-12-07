package patient

import (
	"HMS-16-BE/model"
	"HMS-16-BE/usecase/patient"
	"HMS-16-BE/util/uuid"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type guardianController struct {
	guardian patient.GuardianUsecase
}

func NewGuardianController(g patient.GuardianUsecase) *guardianController {
	return &guardianController{g}
}

func (g *guardianController) GetById(c echo.Context) error {
	id := c.Param("id")

	guardian, err := g.guardian.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    guardian,
	})
}

func (g *guardianController) Create(c echo.Context) error {
	patientId := c.Param("id")
	var guardian model.Guardians
	c.Bind(&guardian)
	guardian.PatientId = patientId
	guardian.Id = uuid.CreateUUID()
	guardian.CreatedAt = time.Now()
	guardian.UpdatedAt = guardian.CreatedAt

	validate := validator.New()
	err := validate.Struct(&guardian)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = g.guardian.Create(guardian)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}

func (g *guardianController) Update(c echo.Context) error {
	id := c.Param("id")
	var guardian model.Guardians
	guardian.Id = id
	guardian.UpdatedAt = time.Now()
	c.Bind(&guardian)

	validate := validator.New()
	err := validate.Struct(&guardian)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = g.guardian.Update(guardian)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}

func (g *guardianController) Delete(c echo.Context) error {
	id := c.Param("id")

	err := g.guardian.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}
