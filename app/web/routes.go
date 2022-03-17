package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tieppt/gob/app/web/controllers"
	"github.com/tieppt/gob/app/web/middlewares"
	"go.uber.org/zap"
)

func SetupRoutes(app *fiber.App, logger *zap.Logger) {
	app.Get("/", welcome)
	routes := app.Group("/api/v1")
	// Public routes
	routes.Post("register", controllers.Register)
	routes.Post("login", controllers.Login(logger))

	// JWT Middleware
	app.Use(middlewares.NewJWTMiddleware(logger))
	// Protected routes
	routes.Get("user", controllers.GetCurrentUser)
}

func welcome(ctx *fiber.Ctx) error {
	return ctx.SendString("Welcome to GoB")
}
