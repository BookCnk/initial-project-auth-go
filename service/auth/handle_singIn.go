package service

import (
	"github.com/gofiber/fiber/v2"
	"initial-project-go/entity"
	"initial-project-go/util"
)

func (a authService) HandleSingIn(c *fiber.Ctx) error {
	var request entity.SignInRequest

	err := util.ValidateRequest(c, &request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userEnt, err := a.userRepository.CreateUserIfNotExists(request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	userToken, err := a.userTokenRepository.GenerateToken(userEnt)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	secret, err := a.encryptorRepository.GetPassphrase()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	token, err := userToken.ToToken(secret.Hash)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"email": a.encryptorRepository.Decrypt(userEnt.Email),
		"exp":   userToken.Exp,
		"iat":   userToken.Iat,
		"name":  a.encryptorRepository.Decrypt(userEnt.Name),
		"token": token,
	})
}
