package shift

import (
	"HMS-16-BE/model"
	"HMS-16-BE/usecase/shift"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type dayController struct {
	day shift.DayUsecase
}

func NewDayController(d shift.DayUsecase) *dayController {
	return &dayController{d}
}

func (t *dayController) GetAll(c echo.Context) error {
	days, err := t.day.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    days,
	})
}

func (t *dayController) GetById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	day, err := t.day.GetById(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    day,
	})
}

func (t *dayController) Create(c echo.Context) error {
	var day model.Days
	c.Bind(&day)

	validate := validator.New()
	err := validate.Struct(&day)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = t.day.Create(day)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}

func (t *dayController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var day model.Days
	c.Bind(&day)
	day.ID = uint(id)

	err := t.day.Update(day)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}

func (t *dayController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	err = t.day.Delete(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}
