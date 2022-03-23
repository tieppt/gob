package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/zap"
)

func SetupCORS(app *fiber.App, logger *zap.Logger) {
	app.Use(cors.New(cors.Config{
		AllowHeaders: "Content-Type, Authorization",
		AllowOrigins: "*", // we can specify origins here
	}))
	logger.Info("CORS middleware enabled")
}
