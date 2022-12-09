package schedule

import (
	"HMS-16-BE/model"
	"HMS-16-BE/usecase/schedule"
	"HMS-16-BE/util/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type shiftController struct {
	shift schedule.ShiftUsecase
}

func NewShiftController(s schedule.ShiftUsecase) *shiftController {
	return &shiftController{s}
}

func (s *shiftController) GetAllByUserId(c echo.Context) error {
	id := middleware.GetIdJWT(c)

	shifts, err := s.shift.GetAllByUserId(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    shifts,
	})
}

func (s *shiftController) GetById(c echo.Context) error {
	id := c.Param("id")

	shift, err := s.shift.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    shift,
	})
}

func (s *shiftController) Create(c echo.Context) error {
	var shift model.Shifts
	c.Bind(&shift)

	validate := validator.New()
	err := validate.Struct(shift)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = s.shift.Create(shift)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}

func (s *shiftController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var shift model.Shifts
	c.Bind(&shift)
	shift.ID = uint(id)

	err := s.shift.Update(shift)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}

func (s *shiftController) Delete(c echo.Context) error {
	id := c.Param("id")

	err := s.shift.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}
