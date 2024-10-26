package service

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"initial-project-go/di/config"
	"initial-project-go/repository"
)

type (
	AuthService interface {
		HandleSingIn(c *fiber.Ctx) error
		HandleLogin(c *fiber.Ctx) error
		HandlerMe(c *fiber.Ctx) error
	}

	authService struct {
		db                  *gorm.DB
		config              config.AppConfig
		userRepository      repository.UserRepository
		userTokenRepository repository.UserTokenRepository
		encryptorRepository repository.EncryptorRepository
		userAuthRepository  repository.UserAuthRepository
		roleRepository      repository.RoleRepository
	}
)

func ProvideAuthService(db *gorm.DB, config config.AppConfig) AuthService {
	return &authService{
		db:                  db,
		config:              config,
		userRepository:      repository.ProvideUserRepository(db, config),
		userTokenRepository: repository.ProvideUserTokenRepository(db, config),
		encryptorRepository: repository.ProvideEncryptorRepository(db, config),
		userAuthRepository:  repository.ProvideUserAuthRepository(db, config),
		roleRepository:      repository.ProvideRoleRepository(db, config),
	}
}
