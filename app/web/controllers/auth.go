package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tieppt/gob/app/business/dto"
	"github.com/tieppt/gob/app/business/services"
	"go.uber.org/zap"
	"net/http"
)

func Register(c *fiber.Ctx) error {
	var user dto.UserRegisterDto
	if err := c.BodyParser(&user); err != nil {
		response := dto.GenericErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid or missing data",
		}
		return c.Status(response.Status).JSON(response)
	}

	result, err := services.CreateUser(&user)

	if err != nil {
		response := dto.GenericErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Can not create user",
		}
		return c.Status(response.Status).JSON(response)
	}

	return c.JSON(result)
}

func Login(logger *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user dto.UserLoginDto
		if err := c.BodyParser(&user); err != nil {
			response := dto.GenericErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid or missing data",
			}
			logger.Error("Invalid or missing data", zap.Error(err))
			return c.Status(response.Status).JSON(response)
		}

		result, err := services.Login(&user)

		if err != nil {
			response := dto.GenericErrorResponse{
				Status:  http.StatusUnauthorized,
				Message: "Wrong username or password",
			}
			logger.Error("Wrong username or password", zap.Error(err))
			return c.Status(response.Status).JSON(response)
		}

		return c.JSON(result)
	}
}
