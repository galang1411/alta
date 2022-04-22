package controller

import (
	"app-echo/domain/model"
	"app-echo/domain/repository"
	"fmt"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	repo repository.UserRepository
}

func NewUserController(repo repository.UserRepository) UserController {
	return UserController{
		repo: repo,
	}
}

func (uc UserController) GetAll(c echo.Context) error {
	users, err := uc.repo.GetAll()
	if err != nil {
		fmt.Println(err)
		return c.JSON(500, echo.Map{
			"error": "cannot fetch user",
		})
	}
	return c.JSON(200, users)
}

func (uc UserController) Create(c echo.Context) error {
	newUser := model.User{}
	c.Bind(&newUser)

	user, err := uc.repo.Create(newUser)
	if err != nil {
		fmt.Print(err.Error())
		return c.JSON(500, echo.Map{
			"error": "cannot add user",
		})
	}
	return c.JSON(201, user)
}
