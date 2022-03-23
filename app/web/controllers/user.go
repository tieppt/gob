package controllers

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/tieppt/gob/app/business/dto"
	"github.com/tieppt/gob/app/business/services"
	"gorm.io/gorm"
)

func GetCurrentUser(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["iss"].(string)

	user, err := services.GetUserById(userId)

	if err != nil {
		response := dto.GenericErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Something went wrong",
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response = dto.GenericErrorResponse{
				Status:  http.StatusNotFound,
				Message: "User not found",
			}
		}
		return c.Status(response.Status).JSON(response)
	}

	return c.JSON(user)
}
