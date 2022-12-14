package shift

import (
	"HMS-16-BE/model"
	"HMS-16-BE/usecase/shift"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type shiftController struct {
	shift shift.ShiftUsecase
}

func NewShiftController(s shift.ShiftUsecase) *shiftController {
	return &shiftController{s}
}

func (s *shiftController) GetAll(c echo.Context) error {
	shifts, err := s.shift.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    shifts,
	})
}

func (s *shiftController) GetAllByUserId(c echo.Context) error {
	userId := c.Param("id")

	shifts, err := s.shift.GetAllByUserId(userId)
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
	err := validate.Struct(&shift)
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
