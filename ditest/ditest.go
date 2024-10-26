package ditest

import (
	config2 "initial-project-go/di/config"
	databasetest "initial-project-go/ditest/database"
	"initial-project-go/repository"
	service "initial-project-go/service/api_keys"
	"testing"
)

func InitTestApplication(t *testing.T) *ApplicationTestSuite {
	db, cleanup := databasetest.InitTestDatabase()
	config := config2.GetConfig()
	encryptorRepository := repository.ProvideEncryptorRepository(db, config)
	_, err := encryptorRepository.GeneratePassphrase(32)
	if err != nil {
		panic(err)
	}

	applicationTestSuite := &ApplicationTestSuite{
		DB:                  db,
		Config:              config,
		ApiKeysRepository:   repository.ProvideApiKeysRepository(db, config),
		EncryptorRepository: encryptorRepository,
		ApiKeysService:      service.ProvideApiKeysService(db, config),
	}

	t.Cleanup(func() {
		cleanup()
	})

	return applicationTestSuite
}
