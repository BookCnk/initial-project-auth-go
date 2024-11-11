package repository

import (
	"gorm.io/gorm"
	"initial-project-go/di/config"
	"initial-project-go/entity"
)

type (
	ProjectRepository interface {
		GetAllProject() ([]entity.Project, error)
		CreateProject(req entity.CreateProjectReq) (entity.Project, error)
	}

	projectRepository struct {
		db                  *gorm.DB
		config              config.AppConfig
		encryptorRepository EncryptorRepository
		userAuthRepository  UserAuthRepository
	}
)

func (p projectRepository) GetAllProject() ([]entity.Project, error) {
	var projects []entity.Project
	if err := p.db.Find(&projects).Error; err != nil {
		return []entity.Project{}, err
	}
	return projects, nil
}

func (p projectRepository) CreateProject(req entity.CreateProjectReq) (entity.Project, error) {
	project := entity.Project{
		Name:        req.Name,
		Description: req.Description,
	}
	if err := p.db.Create(&project).Error; err != nil {
		return entity.Project{}, err
	}
	return project, nil
}

func ProvideProjectRepository(db *gorm.DB, config config.AppConfig) ProjectRepository {
	return &projectRepository{
		db:                  db,
		config:              config,
		encryptorRepository: ProvideEncryptorRepository(db, config),
		userAuthRepository:  ProvideUserAuthRepository(db, config),
	}
}
