package controller

import (
	"HMS-16-BE/model"
	"HMS-16-BE/usecase/admin"
	"HMS-16-BE/util/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func NewAdminController(a usecase.AdminUsecase) *adminController {
	return &adminController{a}
}

type adminController struct {
	admin usecase.AdminUsecase
}

func (a *adminController) Create(c echo.Context) error {
	var admin model.Admins
	c.Bind(&admin)
	admin.ID = util.CreateUUID()
	admin.CreatedAt = time.Now()
	admin.UpdatedAt = admin.CreatedAt

	err := a.admin.Create(admin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success create",
	})
}

func (a *adminController) Login(c echo.Context) error {
	var admin model.Admins
	c.Bind(&admin)

	DTOAdmin, err := a.admin.Login(admin.Username, admin.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success login",
		"data":    DTOAdmin,
	})
}

func (a *adminController) GetById(c echo.Context) error {
	id := c.Param("id")

	admin, err := a.admin.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get",
		"data":    admin,
	})
}

func (a *adminController) Update(c echo.Context) error {
	var admin model.Admins
	c.Bind(&admin)
	admin.UpdatedAt = time.Now()

	id := c.Param("id")
	admin.ID = id

	err := a.admin.Update(admin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update",
	})
}

func (a *adminController) Delete(c echo.Context) error {
	id := c.Param("id")

	err := a.admin.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete",
	})
}
