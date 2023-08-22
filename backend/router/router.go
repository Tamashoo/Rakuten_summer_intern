package router

import (
	"backend/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController, rcc controller.IReceiptController,
	rdc controller.IReduceController, cc controller.ICharacterController) *echo.Echo {

	e := echo.New()
	e.Use(middleware.CORS())
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.LogOut)
	e.GET("/home", cc.Home)
	e.POST("/receipt", rcc.GetReceipt)
	e.GET("/receiptresult", rcc.GetReceiptResult)
	e.GET("/history", rcc.GetHistory)
	e.GET("/foodlossreduce/person", rdc.GetUserReduce)
	e.GET("/foodlossreduce/all", rdc.GetAllReduce)
	return e
}
