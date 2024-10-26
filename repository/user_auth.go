package repository

import (
	"gorm.io/gorm"
	"initial-project-go/di/config"
	"initial-project-go/entity"
)

type UserAuthRepository interface {
	CreateUserAuth(user entity.User, password string) (entity.UserAuth, error)
	FindByUserID(UserID string) (entity.UserAuth, error)
}

type userAuthRepository struct {
	db                  *gorm.DB
	config              config.AppConfig
	encryptorRepository EncryptorRepository
}

func (u userAuthRepository) FindByUserID(UserID string) (entity.UserAuth, error) {
	var userAuth entity.UserAuth
	if err := u.db.Where("user_id = ?", UserID).First(&userAuth).Error; err != nil {
		return entity.UserAuth{}, err
	}
	return userAuth, nil
}

func (u userAuthRepository) CreateUserAuth(user entity.User, password string) (entity.UserAuth, error) {
	id, err := u.encryptorRepository.GeneratePassphrase(30)
	if err != nil {
		return entity.UserAuth{}, err
	}

	salt, err := u.encryptorRepository.GeneratePassphrase(10)
	if err != nil {
		return entity.UserAuth{}, err
	}

	hashesPassword := u.encryptorRepository.HashPassword(salt, password)

	userAuthEnt := entity.UserAuth{
		ID:     id,
		UserID: user.ID,
		Salt:   salt,
		Hash:   hashesPassword,
	}

	result := u.db.Create(&userAuthEnt)
	if result.Error != nil {
		return entity.UserAuth{}, result.Error
	}
	return userAuthEnt, nil
}

func ProvideUserAuthRepository(db *gorm.DB, config config.AppConfig) UserAuthRepository {
	return &userAuthRepository{
		db:                  db,
		config:              config,
		encryptorRepository: ProvideEncryptorRepository(db, config),
	}
}
