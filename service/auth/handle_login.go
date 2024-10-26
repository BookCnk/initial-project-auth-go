package service

import (
	"github.com/gofiber/fiber/v2"
	"initial-project-go/entity"
	"initial-project-go/util"
)

func (a authService) HandleLogin(c *fiber.Ctx) error {
	var request entity.LoginRequest

	err := util.ValidateRequest(c, &request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"err": err.Error(),
		})
	}

	user, err := a.userRepository.FindByEmail(request.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"err": err.Error(),
		})
	}

	userAuth, err := a.userAuthRepository.FindByUserID(user.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"err": err.Error(),
		})
	}

	hashedInputPassword := a.encryptorRepository.HashPassword(userAuth.Salt, request.Password)

	if hashedInputPassword != userAuth.Hash {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"err": "invalid Password",
		})
	}

	token, err := a.userTokenRepository.GenerateToken(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": err.Error(),
		})
	}

	secretKey, err := a.encryptorRepository.GetPassphrase()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": err.Error(),
		})
	}

	tokenJWT, err := token.ToToken(secretKey.Hash)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login Success",
		"token":   tokenJWT,
	})

}
