package controllers

import (
	"CICD_GolangtoHeroku/models"
	"CICD_GolangtoHeroku/services"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	CreateUser(c echo.Context) error
	GetAllUsers(c echo.Context) error
	LoginUser(c echo.Context) error
}

type userController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *userController {
	return &userController{userService}
}

func (u *userController) GetAllUsers(c echo.Context) error {
	users, err := u.userService.GetAllUsers()
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(200, echo.Map{
		"data": users,
	})
}

func (u *userController) CreateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}

	user, err := u.userService.CreateUser(user)
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(200, echo.Map{
		"data": user,
	})
}

func (u *userController) LoginUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}

	userResponse, err := u.userService.LoginUser(user.Email, user.Password)
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(200, echo.Map{
		"data": userResponse,
	})
}
