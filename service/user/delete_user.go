package service

import "github.com/gofiber/fiber/v2"

func (u userService) HandelDeleteUser(c *fiber.Ctx) error {
	userID := c.Params("id")

	err := u.userRepository.DeleteUser(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User delete successfully"})

}
