package middlewares

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/tieppt/gob/app/business/dto"
	"github.com/tieppt/gob/app/foundation/config"
	"go.uber.org/zap"
)

func NewJWTMiddleware(logger *zap.Logger) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   config.SigningKey,
		ErrorHandler: jwtErrorHandler(logger),
	})
}

func jwtErrorHandler(logger *zap.Logger) fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		logger.Error("JWT Error", zap.Error(err))
		response := dto.GenericErrorResponse{
			Status:  401,
			Message: err.Error(),
		}
		return c.Status(response.Status).JSON(response)
	}
}
