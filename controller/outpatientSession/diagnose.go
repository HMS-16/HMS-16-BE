package outpatientSession

import (
	"HMS-16-BE/model"
	"HMS-16-BE/usecase/outpatientSession"
	"HMS-16-BE/util/middleware"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

type diagnoseController struct {
	diagnose outpatientSession.DiagnoseUsecase
}

func NewDiagnoseController(d outpatientSession.DiagnoseUsecase) *diagnoseController {
	return &diagnoseController{d}
}

func (d *diagnoseController) Create(c echo.Context) error {
	patientId := c.Param("id")
	var diagnose model.Diagnoses
	c.Bind(&diagnose)
	diagnose.DoctorId = middleware.GetIdJWT(c)
	diagnose.CreatedAt = time.Now()
	diagnose.UpdatedAt = diagnose.CreatedAt

	err := d.diagnose.Create(diagnose, patientId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}

func (d *diagnoseController) GetById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	diagnose, err := d.diagnose.GetById(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    diagnose,
	})
}

func (d *diagnoseController) GetAllByPatient(c echo.Context) error {
	id := c.Param("id")

	diagnose, err := d.diagnose.GetAllByPatient(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message":    "success",
		"patient_id": id,
		"data":       diagnose,
	})
}