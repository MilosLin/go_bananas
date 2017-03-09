package router

import (
	"github.com/MilosLin/go_bananas/api/controller"

	validator "gopkg.in/go-playground/validator.v9"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

/**
 * 初始化路由
 */
func InitRouting(server *echo.Echo) {
	v := validator.New()
	server.Validator = &CustomValidator{validator: v}
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.Use(middleware.CORS())
	server.POST("/api/SetGameConfig", controller.PostSetGameConfig)
	server.GET("/api/wagers", controller.GetWagers)
	server.GET("/api/wager/:game_seq_no", controller.GetWager)
	server.GET("/api/ccu", controller.GetCCU)
	server.GET("/api/financialreport", controller.GetFinancialReport)
	server.GET("/api/supervisorreport", controller.GetSupervisorReport)
	server.GET("/api/examinewager", controller.GetExamineWager)
	server.GET("/api/abnormalwager", controller.GetAbnormalWager)
}
