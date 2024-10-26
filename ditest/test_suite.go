package ditest

import (
	"gorm.io/gorm"
	"initial-project-go/di/config"
	"initial-project-go/repository"
	"initial-project-go/service/api_keys"
)

type ApplicationTestSuite struct {
	DB                  *gorm.DB
	Config              config.AppConfig
	ApiKeysRepository   repository.ApiKeysRepository
	EncryptorRepository repository.EncryptorRepository
	ApiKeysService      service.ApiKeysService
}
