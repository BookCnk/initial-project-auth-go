package service

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"initial-project-go/di/config"
	"initial-project-go/repository"
)

type ApiKeysService interface {
	HandleSecretPostRouter(c *fiber.Ctx) error
	HandleRotatePostRouter(c *fiber.Ctx) error
	HandleVerifyGetRouter(c *fiber.Ctx) error
}

type apiKeysService struct {
	apiKeysRepository   repository.ApiKeysRepository
	encryptorRepository repository.EncryptorRepository
	db                  *gorm.DB
	config              config.AppConfig
}

func ProvideApiKeysService(db *gorm.DB, config config.AppConfig) ApiKeysService {
	apiKeysRepository := repository.ProvideApiKeysRepository(db, config)
	encryptorRepository := repository.ProvideEncryptorRepository(db, config)
	return &apiKeysService{
		db:                  db,
		config:              config,
		apiKeysRepository:   apiKeysRepository,
		encryptorRepository: encryptorRepository,
	}
}
