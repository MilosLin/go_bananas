package api

import (
	"github.com/MilosLin/go_bananas/core/config"
	"github.com/MilosLin/go_bananas/core/model/database"
	"github.com/api/router"

	"github.com/labstack/echo"
)

type API_Service struct {
}

func NewAPIService() *API_Service {
	return &API_Service{}
}

/**
 * 啟動API服務
 */
func (s *API_Service) Start() {
	defer database.CloseConn()
	server := echo.New()
	router.InitRouting(server)
	server.Logger.Fatal(server.Start(":" + config.Instance().GetString("api.listen_port")))
}
