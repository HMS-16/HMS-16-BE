package outpatientSession

import (
	"HMS-16-BE/model"
	"HMS-16-BE/usecase/outpatientSession"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type conditionController struct {
	condition outpatientSession.ConditionUsecase
}

func NewConditionController(c outpatientSession.ConditionUsecase) *conditionController {
	return &conditionController{c}
}

func (d *conditionController) Create(c echo.Context) error {
	var condition model.Conditions
	c.Bind(&condition)

	validate := validator.New()
	err := validate.Struct(&condition)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = d.condition.Create(condition)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}

func (d *conditionController) GetById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	condition, err := d.condition.GetById(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    condition,
	})
}