package shift

import (
	"HMS-16-BE/model"
	"HMS-16-BE/usecase/shift"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type timeController struct {
	time shift.TimeUsecase
}

func NewTimeController(t shift.TimeUsecase) *timeController {
	return &timeController{t}
}

func (t *timeController) GetAll(c echo.Context) error {
	times, err := t.time.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    times,
	})
}

func (t *timeController) GetById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	time, err := t.time.GetById(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    time,
	})
}

func (t *timeController) Create(c echo.Context) error {
	var time model.Times
	c.Bind(&time)

	validate := validator.New()
	err := validate.Struct(&time)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = t.time.Create(time)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}

func (t *timeController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var time model.Times
	c.Bind(&time)
	time.ID = uint(id)

	err := t.time.Update(time)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}

func (t *timeController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	err = t.time.Delete(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}
