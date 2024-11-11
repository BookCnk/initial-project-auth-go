package project

import "github.com/gofiber/fiber/v2"

func (p projectService) HandlerGetAllProject(c *fiber.Ctx) error {
	result, err := p.projectRepository.GetAllProject()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}
