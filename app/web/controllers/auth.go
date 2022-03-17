package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tieppt/gob/app/business/dto"
	"github.com/tieppt/gob/app/business/services"
	"go.uber.org/zap"
)

func Register(c *fiber.Ctx) error {
	var user dto.UserRegisterDto
	if err := c.BodyParser(&user); err != nil {
		response := dto.GenericErrorResponse{
			Status:  400,
			Message: "Invalid or missing data",
		}
		return c.Status(response.Status).JSON(response)
	}

	result, err := services.CreateUser(&user)

	if err != nil {
		response := dto.GenericErrorResponse{
			Status:  500,
			Message: "Can not create user",
		}
		return c.Status(response.Status).JSON(response)
	}

	return c.JSON(result)
}

func Login(logger *zap.Logger) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var user dto.UserLoginDto
		if err := c.BodyParser(&user); err != nil {
			response := dto.GenericErrorResponse{
				Status:  400,
				Message: "Invalid or missing data",
			}
			return c.Status(response.Status).JSON(response)
		}

		result, err := services.Login(&user)

		if err != nil {
			response := dto.GenericErrorResponse{
				Status:  401,
				Message: "Wrong username or password",
			}
			return c.Status(response.Status).JSON(response)
		}

		return c.JSON(result)
	}
}
