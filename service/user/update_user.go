package service

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"initial-project-go/entity"
)

func (u userService) HandleUpdateUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	if userID == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	var updateUserReq entity.UpdateUserReq
	if err := c.BodyParser(&updateUserReq); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	updatedUser, err := u.userRepository.UpdateUser(userID, updateUserReq)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(updatedUser)
}
