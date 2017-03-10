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
	GetPingOutput struct {
		IP       string `json:"ip"`
		DateTime string `json:"datetime"`
	}
)

/**
 * 取得單一注單
 * router /api/ccu
 */
func GetPing(c echo.Context) error {
	res := protocol.Response{}
	output := GetPingOutput{}
	output.DateTime = time.Now().Format(env.DateTimeWithTimeZone)
	output.IP = cell.GetIPv4().String()

	res.Code = protocol.E_OK
	res.Data = output

	return c.JSON(http.StatusOK, res)
}
