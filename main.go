package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tieppt/gob/app/foundation/config"
	"github.com/tieppt/gob/app/foundation/database"
	"github.com/tieppt/gob/app/web"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	// load config
	config.InitConfig(logger)

	database.ConnectDB(logger)

	app := fiber.New()

	web.SetupRoutes(app, logger)
	web.SetupCORS(app, logger)

	sugar.Infof("start app at: %v", config.Conf.Port)
	sugar.Error(app.Listen(":" + config.Conf.Port))
}
