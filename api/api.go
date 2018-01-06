package api

import (
	"github.com/MilosLin/go_bananas/api/router"
	"github.com/MilosLin/go_bananas/core/config"
	"github.com/MilosLin/go_bananas/core/model/database"

	"github.com/labstack/echo"
)

// APIService : API Service
type APIService struct {
}

// NewAPIService : New API Service Instance
func NewAPIService() *APIService {
	return &APIService{}
}

// Start : 啟動API服務
func (s *APIService) Start() {
	defer database.CloseConn()
	server := echo.New()
	router.InitRouting(server)
	server.Start(":" + config.Instance().GetString("api.listenPort"))
}
