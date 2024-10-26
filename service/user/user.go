package service

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"initial-project-go/di/config"
	"initial-project-go/repository"
)

type UserService interface {
	HandleGetAllUser(c *fiber.Ctx) error
	HandelCreatUser(c *fiber.Ctx) error
	HandelDeleteUser(c *fiber.Ctx) error
	HandleHarddel(c *fiber.Ctx) error
	//HandelCreatUser(c *fiber.Ctx) error
	HandleUpdateUser(c *fiber.Ctx) error
}

type userService struct {
	db             *gorm.DB
	config         config.AppConfig
	userRepository repository.UserRepository
}

func ProvideUserService(db *gorm.DB, config config.AppConfig) UserService {
	return &userService{
		db:             db,
		config:         config,
		userRepository: repository.ProvideUserRepository(db, config),
	}
}
