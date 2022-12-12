package user

import (
	"HMS-16-BE/dto"
	"HMS-16-BE/model"
	"HMS-16-BE/usecase/user"
	"HMS-16-BE/util/hash"
	"HMS-16-BE/util/middleware"
	"HMS-16-BE/util/uuid"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type userController struct {
	user user.UserUsecase
}

func NewUserController(u user.UserUsecase) *userController {
	return &userController{u}
}

func (u *userController) Create(c echo.Context) error {
	var user model.Users
	c.Bind(&user)
	user.Id = uuid.CreateUUID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
	user.Password, _ = hash.HashPassword(user.Password)

	validate := validator.New()
	err := validate.Struct(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = u.user.Create(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success create",
	})
}

func (u *userController) Login(c echo.Context) error {
	var userInput model.Users
	c.Bind(&userInput)

	user, err := u.user.Login(userInput.Email)
	if err != nil {
		return c.JSON(http.StatusForbidden, err.Error())
	}

	if !hash.CheckPasswordHash(userInput.Password, user.Password) {
		return c.JSON(http.StatusForbidden, err.Error())
	}

	token, _ := middleware.CreateToken(user.Id, user.Email, dto.Role[user.Role])

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    *dto.UserDTO(&user),
		"token":   token,
	})
}

func (u *userController) GetAll(c echo.Context) error {
	users, err := u.user.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    users,
	})
}

func (u *userController) GetById(c echo.Context) error {
	id := c.Param("id")

	user, err := u.user.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    user,
	})
}

func (u *userController) Update(c echo.Context) error {
	var user model.Users
	c.Bind(&user)
	user.UpdatedAt = time.Now()

	id := c.Param("id")
	user.Id = id

	err := u.user.Update(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}

func (u *userController) Delete(c echo.Context) error {
	id := c.Param("id")

	err := u.user.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}
