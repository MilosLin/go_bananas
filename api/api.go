package api

import (
	"github.com/MilosLin/go_bananas/api/router"
	"github.com/MilosLin/go_bananas/core/config"
	"github.com/MilosLin/go_bananas/core/model/database"

	"github.com/labstack/echo"
)

type APIService struct {
}

func NewAPIService() *APIService {
	return &APIService{}
}

/**
 * 啟動API服務
 */
func (s *APIService) Start() {
	defer database.CloseConn()
	server := echo.New()
	router.InitRouting(server)
	server.Start(":" + config.Instance().GetString("api.listenPort"))
}
