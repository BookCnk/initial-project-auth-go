package service

import (
	"github.com/gofiber/fiber/v2"
)

func (u userService) HandleGetAllUser(c *fiber.Ctx) error {
	users, err := u.userRepository.GetAllUser()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})

	}
	return c.Status(fiber.StatusOK).JSON(users)
}
