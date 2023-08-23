package controller

import (
	"backend/model"
	"backend/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IReduceController interface {
	GetUserReduce(c echo.Context) error
	GetAllReduce(c echo.Context) error
}

type reduceController struct {
	ru usecase.IReduceUsecase
}

func NewReduceController(ru usecase.IReduceUsecase) IReduceController {
	return &reduceController{ru}
}

func (rc *reduceController) GetUserReduce(c echo.Context) error {
	cookie, err := c.Cookie("username")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.UserReduceResponse{})
	}
	username := cookie.Value
	urrRes := model.UserReduceResponse{}
	score, err := rc.ru.GetUserReduceByUsername(username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, urrRes)
	}
	urrRes.ReduceScore = score
	return c.JSON(http.StatusOK, urrRes)
}

func (rc *reduceController) GetAllReduce(c echo.Context) error {
	arr := model.AllReduceResponse{}
	score, err := rc.ru.GetAllReduce()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, arr)
	}
	arr.AllReduce = score
	return c.JSON(http.StatusOK, arr)
}
