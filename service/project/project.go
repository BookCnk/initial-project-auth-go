package project

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"initial-project-go/di/config"
	"initial-project-go/repository"
)

type (
	ProjectService interface {
		HandlerCreateProject(c *fiber.Ctx) error
		HandlerGetAllProject(c *fiber.Ctx) error
	}

	projectService struct {
		db                *gorm.DB
		config            config.AppConfig
		projectRepository repository.ProjectRepository
	}
)

func ProvideProjectService(db *gorm.DB, config config.AppConfig) ProjectService {
	return &projectService{
		db:                db,
		config:            config,
		projectRepository: repository.ProvideProjectRepository(db, config),
	}
}
