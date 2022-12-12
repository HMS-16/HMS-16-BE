package profile

import (
	"HMS-16-BE/model"
	"HMS-16-BE/usecase/profile"
	"HMS-16-BE/util/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type nurseController struct {
	nurse profile.NurseUsecase
}

func NewNurseController(n profile.NurseUsecase) *nurseController {
	return &nurseController{n}
}

func (n *nurseController) GetAll(c echo.Context) error {
	nurses, err := n.nurse.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    nurses,
	})
}

func (n *nurseController) GetById(c echo.Context) error {
	id := middleware.GetIdJWT(c)

	nurse, err := n.nurse.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    nurse,
	})
}

func (n *nurseController) Create(c echo.Context) error {
	var nurse model.Nurses
	c.Bind(&nurse)
	nurse.CreatedAt = time.Now()
	nurse.UpdatedAt = nurse.CreatedAt
	nurse.UserId = middleware.GetIdJWT(c)

	validate := validator.New()
	err := validate.Struct(&nurse)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = n.nurse.Create(nurse)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}

func (n *nurseController) Update(c echo.Context) error {
	var nurse model.Nurses
	c.Bind(&nurse)
	nurse.UserId = middleware.GetIdJWT(c)
	nurse.UpdatedAt = time.Now()

	validate := validator.New()
	err := validate.Struct(nurse)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = n.nurse.Update(nurse)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}

func (n *nurseController) Delete(c echo.Context) error {
	id := middleware.GetIdJWT(c)

	err := n.nurse.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}
