package handlers

import (
	"shopping-site/api/services"
	"shopping-site/internals"
	"shopping-site/pkg/models"
	"shopping-site/utils/dto"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	services.UserService
}

func (userservice *UserHandler) GetUsersHandler(ctx *fiber.Ctx) error {
	var user []models.Users

	record := internals.Pg.Model(&models.Users{}).Find(&user)
	if record.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ResponseJson{
			Error: record.Error.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.ResponseJson{
		Data: user,
	})

}
