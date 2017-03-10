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
func InitRouting(e *echo.Echo) {
	v := validator.New()
	e.Validator = &CustomValidator{validator: v}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	v1 := e.Group("/v1")
	v1.Use()
	v1.GET("/api/ping", controller.GetPing)

}
