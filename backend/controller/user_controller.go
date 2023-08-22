package controller

import (
	"backend/model"
	"backend/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
	LogOut(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

func (uc *userController) SignUp(c echo.Context) error {
	user := model.User{}
	userRes := model.UserResponse{}
	userRes.Result = false
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, userRes)
	}
	if err := uc.uu.SignUp(user); err != nil {
		return c.JSON(http.StatusInternalServerError, userRes)
	}
	userRes.Result = true
	return c.JSON(http.StatusCreated, userRes)
}

func (uc *userController) Login(c echo.Context) error {
	user := model.User{}
	userRes := model.UserResponse{}
	userRes.Result = false
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, userRes)
	}
	_, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, userRes)
	}
	userRes.Result = true
	return c.JSON(http.StatusOK, userRes)
}

func (uc *userController) LogOut(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
