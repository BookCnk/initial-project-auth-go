package project

import (
	"github.com/gofiber/fiber/v2"
	"initial-project-go/entity"
	"initial-project-go/util"
)

func (p projectService) HandlerCreateProject(c *fiber.Ctx) error {
	var createProjectRequest entity.CreateProjectReq

	err := util.ValidateRequest(c, &createProjectRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	result, err := p.projectRepository.CreateProject(createProjectRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}
