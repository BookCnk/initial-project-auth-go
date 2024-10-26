package service

import (
	"github.com/gofiber/fiber/v2"
	"initial-project-go/entity"
	"initial-project-go/util"
)

func (u userService) HandelCreatUser(c *fiber.Ctx) error {
	var createUserRequest entity.CreateUserReq

	err := util.ValidateRequest(c, &createUserRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	result, err := u.userRepository.CreateUser(createUserRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}
