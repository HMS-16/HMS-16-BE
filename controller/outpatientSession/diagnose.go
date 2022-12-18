package outpatientSession

import (
	"HMS-16-BE/model"
	"HMS-16-BE/usecase/outpatientSession"
	"HMS-16-BE/util/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type diagnoseController struct {
	diagnose outpatientSession.DiagnoseUsecase
}

func NewDiagnoseController(d outpatientSession.DiagnoseUsecase) *diagnoseController {
	return &diagnoseController{d}
}

func (d *diagnoseController) Create(c echo.Context) error {
	var diagnose model.Diagnoses
	c.Bind(&diagnose)
	diagnose.DoctorId = middleware.GetIdJWT(c)

	validate := validator.New()
	err := validate.Struct(&diagnose)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = d.diagnose.Create(diagnose)
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
