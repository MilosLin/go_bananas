package controller

import (
	"net/http"
	"time"

	"github.com/MilosLin/go_bananas/api/protocol"
	"github.com/MilosLin/go_bananas/core/env"
	"github.com/MilosLin/go_bananas/core/module/cell"

	"github.com/labstack/echo"
)

type (
	// GetPingOutput : Get Ping API Output
	GetPingOutput struct {
		IP       string `json:"ip"`
		DateTime string `json:"datetime"`
	}
)

// GetPing : ping this service
// router /api/ping
func GetPing(c echo.Context) error {
	res := protocol.Response{}
	output := GetPingOutput{}
	output.DateTime = time.Now().Format(env.DateTimeWithTimeZone)
	output.IP = cell.GetIPv4().String()

	res.Code = protocol.ErrCodeOK
	res.Data = output

	return c.JSON(http.StatusOK, res)
}
