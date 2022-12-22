package controller

import (
	"HMS-16-BE/dto"
	"HMS-16-BE/model"
	"HMS-16-BE/usecase/admin"
	"HMS-16-BE/util/hash"
	"HMS-16-BE/util/middleware"
	"HMS-16-BE/util/uuid"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func NewAdminController(a admin.AdminUsecase) *adminController {
	return &adminController{a}
}

type adminController struct {
	admin admin.AdminUsecase
}

func (a *adminController) Create(c echo.Context) error {
	var admin model.Admins
	c.Bind(&admin)
	admin.ID = uuid.CreateUUID()
	admin.CreatedAt = time.Now()
	admin.UpdatedAt = admin.CreatedAt
	admin.Password, _ = hash.HashPassword(admin.Password)

	validate := validator.New()
	err := validate.Struct(&admin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = a.admin.Create(admin)
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

	adminDB, err := a.admin.Login(admin.Username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if !hash.CheckPasswordHash(admin.Password, adminDB.Password) {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "password incorrect",
		})
	}

	role := "admin"
	token, _ := middleware.CreateToken(adminDB.ID, adminDB.Username, role)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success login",
		"data":    *dto.AdminDTO(&adminDB),
		"token":   token,
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
