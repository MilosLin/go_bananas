package task

import (
	"github.com/MilosLin/go_bananas/core/logger"
	"go.uber.org/zap"
)

type LogExample struct{}

func NewLogExample() *LogExample {
	return &LogExample{}
}

// run task TaskTemplate
//
// task --name=logexample
func (instance *LogExample) Run(argu *string) error {
	// write default log
	logger.Debug("debug", zap.String("key1", "some value"))
	logger.Info("Info")
	logger.Warn("Warn")
	logger.Error("Error", zap.Stack(""))

	// write critical log
	logger.Forge("critical").Debug("debug")
	logger.Forge("critical").Info("Info")
	logger.Forge("critical").Warn("Warn")
	logger.Forge("critical").Error("Error", zap.Stack(""))
	return nil
}
