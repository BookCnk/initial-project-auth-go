package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"initial-project-go/di/config"
	"initial-project-go/entity"
)

type UserRepository interface {
	GetAllUser() ([]entity.ResponseAllUser, error)
	CreateUser(user entity.CreateUserReq) (entity.User, error)
	DeleteUser(id string) error
	HardDelUser(id string) error
	CreateUserIfNotExists(request entity.SignInRequest) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	FindByID(id string) (entity.User, error)
	UpdateUser(id string, user entity.UpdateUserReq) (entity.User, error)
}

type userRepository struct {
	db                  *gorm.DB
	config              config.AppConfig
	encryptorRepository EncryptorRepository
	userAuthRepository  UserAuthRepository
}

func (u userRepository) UpdateUser(id string, user entity.UpdateUserReq) (entity.User, error) {
	var exitingUser entity.User
	result := u.db.First(&exitingUser, "id = ?", id)
	if result.Error != nil {
		return entity.User{}, result.Error
	}

	if user.Name != "" {
		exitingUser.Name = u.encryptorRepository.Encrypt(user.Name)
	}
	if user.Email != "" {
		exitingUser.Email = u.encryptorRepository.Encrypt(user.Email)
	}
	saveResult := u.db.Save(&exitingUser)
	if saveResult.Error != nil {
		return entity.User{}, saveResult.Error
	}

	return exitingUser, nil
}

func (u userRepository) FindByID(id string) (entity.User, error) {
	var ent entity.User
	result := u.db.First(&ent, "id = ?", id)
	if result.Error != nil {
		return entity.User{}, result.Error
	}

	return ent, nil
}

func (u userRepository) FindByEmail(email string) (entity.User, error) {
	var ent entity.User
	result := u.db.First(&ent, "email = ?", u.encryptorRepository.Encrypt(email))
	if result.Error != nil {
		return entity.User{}, result.Error
	}

	return ent, nil
}

func (u userRepository) HardDelUser(id string) error {
	result := u.db.Unscoped().Delete(&entity.User{}, "id = ?", id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no user found with ID %s", id)
	}

	return nil

}

func (u userRepository) GetAllUser() ([]entity.ResponseAllUser, error) {
	var users []entity.User

	result := u.db.Preload("Role").Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	var userDTOs []entity.ResponseAllUser

	for _, user := range users {
		userDTOs = append(userDTOs, entity.ResponseAllUser{
			Email:    u.encryptorRepository.Decrypt(user.Email),
			Name:     u.encryptorRepository.Decrypt(user.Name),
			Id:       user.ID,
			RoleName: user.Role.Name,
		})
	}

	return userDTOs, nil
}

func (u userRepository) DeleteUser(userID string) error {
	result := u.db.Delete(&entity.User{}, "id = ?", userID)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no user found with ID %s", userID)
	}

	return nil
}

func (u userRepository) CreateUser(user entity.CreateUserReq) (entity.User, error) {

	id, err := u.encryptorRepository.GeneratePassphrase(30)
	if err != nil {
		return entity.User{}, err
	}
	userEntity := entity.User{
		ID:    id,
		Name:  u.encryptorRepository.Encrypt(user.Name),
		Email: u.encryptorRepository.Encrypt(user.Email),
	}

	result := u.db.Create(&userEntity)
	if result.Error != nil {
		return entity.User{}, result.Error
	}

	return userEntity, nil
}

func (u userRepository) CreateUserIfNotExists(request entity.SignInRequest) (entity.User, error) {
	existingUser, err := u.FindByEmail(request.Email)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, err
		}
	}

	if existingUser.ID != "" {
		return existingUser, errors.New("user already exists")
	}

	id, err := u.encryptorRepository.GeneratePassphrase(30)
	if err != nil {
		return entity.User{}, err
	}

	userEnt := entity.User{
		ID:     id,
		Name:   u.encryptorRepository.Encrypt(request.Name),
		Email:  u.encryptorRepository.Encrypt(request.Email),
		RoleID: "base-user-role",
	}

	result := u.db.Create(&userEnt)
	if result.Error != nil {
		return entity.User{}, err
	}

	_, err = u.userAuthRepository.CreateUserAuth(userEnt, request.Password)
	if err != nil {
		return entity.User{}, err
	}

	return userEnt, nil
}

func ProvideUserRepository(db *gorm.DB, config config.AppConfig) UserRepository {
	encryptorRepository := ProvideEncryptorRepository(db, config)
	return &userRepository{
		db:                  db,
		config:              config,
		encryptorRepository: encryptorRepository,
		userAuthRepository:  ProvideUserAuthRepository(db, config),
	}
}
