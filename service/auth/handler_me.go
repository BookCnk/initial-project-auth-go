package service

import (
	"github.com/gofiber/fiber/v2"
	"initial-project-go/entity"
)

func (a authService) HandlerMe(c *fiber.Ctx) error {
	user := c.Locals("user").(entity.User)

	response := entity.ResponseMe{
		Email:       a.encryptorRepository.Decrypt(user.Email),
		Id:          user.ID,
		Name:        a.encryptorRepository.Decrypt(user.Name),
		Permissions: a.roleRepository.GetPermissionByRoleId(user.RoleID),
		Role:        user.RoleID,
	}

	return c.JSON(response)
}
