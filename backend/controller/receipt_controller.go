package controller

import (
	"backend/model"
	"backend/usecase"
	"backend/utils"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type IReceiptController interface {
	GetReceipt(c echo.Context) error
	GetReceiptResult(c echo.Context) error
	GetHistory(c echo.Context) error
}

type receiptController struct {
	ru  usecase.IReceiptUsecase
	rdu usecase.IReduceUsecase
}

func NewReceiptController(ru usecase.IReceiptUsecase, rdu usecase.IReduceUsecase) IReceiptController {
	return &receiptController{ru, rdu}
}

func (rc *receiptController) GetReceipt(c echo.Context) error {
	cookie, err := c.Cookie("username")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ReceiptResponse{})
	}
	username := cookie.Value
	rr := model.ReceiptRequest{}
	rrRes := model.ReceiptResponse{}
	rrRes.Result = false
	if err := c.Bind(&rr); err != nil {
		return c.JSON(http.StatusBadRequest, rrRes)
	}

	exp, cnt, err := utils.ForCharacterController(rr.Receipt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, rrRes)
	}
	receipt := model.Receipt{}
	receipt.Username = username
	receipt.GetExp = exp
	receipt.CreatedAt = time.Now()

	if err := rc.ru.CreateReceipt(receipt); err != nil {
		return c.JSON(http.StatusInternalServerError, rrRes)
	}

	if err := rc.rdu.AddUserScore(username, exp, cnt); err != nil {
		return c.JSON(http.StatusInternalServerError, rrRes)
	}

	if err := rc.rdu.AddAllScore(cnt); err != nil {
		return c.JSON(http.StatusInternalServerError, rrRes)
	}

	rrRes.Result = true

	return c.JSON(http.StatusOK, rrRes)
}

func (rc *receiptController) GetReceiptResult(c echo.Context) error {
	cookie, err := c.Cookie("username")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ReceiptResultResponse{})
	}
	username := cookie.Value
	rrr := model.ReceiptResultResponse{}
	rrr, err = rc.ru.GetCreateReceiptResult(username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, rrr)
	}
	return c.JSON(http.StatusOK, rrr)
}

func (rc *receiptController) GetHistory(c echo.Context) error {
	cookie, err := c.Cookie("username")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.HistoryListResponse{})
	}
	username := cookie.Value
	hlr := model.HistoryListResponse{}
	hlr, err = rc.ru.GetHistoryList(username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, hlr)
	}
	return c.JSON(http.StatusOK, hlr)
}
