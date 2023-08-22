package controller

import (
	"backend/model"
	"backend/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IReceiptController interface {
	GetReceipt(c echo.Context) error
	GetReceiptResult(c echo.Context) error
	GetHistory(c echo.Context) error
}

type receiptController struct {
	ru usecase.IReceiptUsecase
}

func NewReceiptController(ru usecase.IReceiptUsecase) IReceiptController {
	return &receiptController{ru}
}

func (rc *receiptController) GetReceipt(c echo.Context) error {
	rr := model.ReceiptRequest{}
	rrRes := model.ReceiptResponse{}
	rrRes.Result = false
	if err := c.Bind(&rr); err != nil {
		return c.JSON(http.StatusBadRequest, rrRes)
	}

}

func (rc *receiptController) GetReceiptResult(c echo.Context) error {
	ur := model.UserRequest{}
	rrr := model.ReceiptResultResponse{}
	if err := c.Bind(&ur); err != nil {
		return c.JSON(http.StatusBadRequest, rrr)
	}
	rrr, err := rc.ru.GetCreateReceiptResult(ur.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, rrr)
	}
	return c.JSON(http.StatusOK, rrr)
}

func (rc *receiptController) GetHistory(c echo.Context) error {
	ur := model.UserRequest{}
	hlr := model.HistoryListResponse{}
	if err := c.Bind(&ur); err != nil {
		return c.JSON(http.StatusBadRequest, hlr)
	}
	hlr, err := rc.ru.GetHistoryList(ur.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, hlr)
	}
	return c.JSON(http.StatusOK, hlr)
}
