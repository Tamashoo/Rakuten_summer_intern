package controller

import (
	"backend/model"
	"backend/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ICharacterController interface {
	Home(c echo.Context) error
}

type characterController struct {
	cu usecase.ICharacterUsecase
	ru usecase.IReduceUsecase
}

func NewCharacterController(cu usecase.ICharacterUsecase, ru usecase.IReduceUsecase) ICharacterController {
	return &characterController{cu, ru}
}

func (cc *characterController) Home(c echo.Context) error {
	cookie, err := c.Cookie("username")
	crRes := model.CharacterResponse{}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, crRes)
	}
	username := cookie.Value
	exp, err := cc.ru.GetUserReduceByUsername(username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, crRes)
	}
	level := 0       // temporary
	exp = exp % 1000 // temporary
	if exp/1000+1 > 4 {
		level = 4
	} else {
		level = exp/1000 + 1
	}
	character, err := cc.cu.GetCharacterByLevel(level)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, crRes)
	}
	crRes.CharacterImage = character
	crRes.Exp = exp
	crRes.Level = level
	return c.JSON(http.StatusOK, crRes)
}
