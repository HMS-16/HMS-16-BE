package outpatientSession

import (
	"HMS-16-BE/model"
	"HMS-16-BE/usecase/outpatientSession"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

type scheduleController struct {
	schedule outpatientSession.ScheduleUsecase
}

func NewScheduleController(s outpatientSession.ScheduleUsecase) *scheduleController {
	return &scheduleController{s}
}

func (s *scheduleController) Create(c echo.Context) error {
	var schedule model.Schedules
	c.Bind(&schedule)
	schedule.CreatedAt = time.Now()
	schedule.UpdatedAt = schedule.CreatedAt

	err := s.schedule.Create(schedule)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}

func (s *scheduleController) GetAllCardByDay(c echo.Context) error {
	date := c.QueryParam("date")
	if date == "" {
		date = time.Now().Format("01/02/2006")
	}

	schedules, err := s.schedule.GetAllCardByDay(date)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    schedules,
	})
}

func (s *scheduleController) GetAllByDay(c echo.Context) error {
	date := c.QueryParam("date")
	if date == "" {
		date = time.Now().Format("01/02//2006")
	}

	schedules, err := s.schedule.GetAllByDay(date)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    schedules,
	})
}

func (s *scheduleController) GetByScheduleId(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	schedule, err := s.schedule.GetByScheduleId(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    schedule,
	})
}

func (s *scheduleController) GetDetailByPatient(c echo.Context) error {
	id := c.Param("id")

	patient, err := s.schedule.GetDetailByPatient(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    patient,
	})
}

func (s *scheduleController) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	var schedule model.Schedules
	c.Bind(&schedule)
	schedule.ID = uint(id)
	schedule.UpdatedAt = time.Now()

	validate := validator.New()
	err = validate.Struct(&schedule)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = s.schedule.Update(schedule)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}

func (s *scheduleController) UpdateDoctor(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	var schedule model.Schedules
	c.Bind(&schedule)
	schedule.ID = uint(id)
	schedule.UpdatedAt = time.Now()

	validate := validator.New()
	err = validate.Struct(&schedule)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = s.schedule.UpdateDoctor(schedule)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}

func (s *scheduleController) UpdateNurse(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	var schedule model.Schedules
	c.Bind(&schedule)
	schedule.ID = uint(id)
	schedule.UpdatedAt = time.Now()

	validate := validator.New()
	err = validate.Struct(&schedule)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = s.schedule.UpdateNurse(schedule)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}

func (s *scheduleController) UpdateStatus(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = s.schedule.UpdateStatus(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}

func (s *scheduleController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = s.schedule.Delete(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}
