package repository

import (
	"errors"
	"gorm.io/gorm"
	"initial-project-go/di/config"
	"initial-project-go/entity"
	"time"
)

type UserTokenRepository interface {
	GenerateToken(userEnt entity.User) (entity.UserToken, error)
}

type userTokenRepository struct {
	db                  *gorm.DB
	config              config.AppConfig
	encryptorRepository EncryptorRepository
}

func (u userTokenRepository) GenerateToken(userEnt entity.User) (entity.UserToken, error) {
	if userEnt.ID == "" {
		return entity.UserToken{}, errors.New("user ID is empty")
	}

	tokenID, err := u.encryptorRepository.GeneratePassphrase(30)
	if err != nil {
		return entity.UserToken{}, err
	}

	userTokenEnt := entity.UserToken{
		ID:     tokenID,
		UserID: userEnt.ID,
		Iat:    time.Now(),
		Exp:    time.Now().Add(time.Minute * 30),
	}

	result := u.db.Create(&userTokenEnt)
	if result.Error != nil {
		return entity.UserToken{}, result.Error
	}

	return userTokenEnt, nil
}

func ProvideUserTokenRepository(db *gorm.DB, config config.AppConfig) UserTokenRepository {
	return &userTokenRepository{
		db:                  db,
		config:              config,
		encryptorRepository: ProvideEncryptorRepository(db, config),
	}
}
