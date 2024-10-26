package repository

import (
	"gorm.io/gorm"
	"initial-project-go/di/config"
	"initial-project-go/entity"
)

type (
	RoleRepository interface {
		GetPermissionByRoleId(id string) []string
	}

	roleRepository struct {
		db     *gorm.DB
		config config.AppConfig
	}
)

func (r roleRepository) GetPermissionByRoleId(id string) []string {

	var roleEnt entity.Role
	err := r.db.Preload("Permissions").First(&roleEnt, "id = ? ", id).Error
	if err != nil {
		return []string{}
	}

	permissionsEnt := roleEnt.Permissions
	var permissions []string

	for _, permission := range permissionsEnt {
		permissions = append(permissions, permission.ID)

	}

	return permissions

}

func ProvideRoleRepository(db *gorm.DB, appConfig config.AppConfig) RoleRepository {
	return &roleRepository{
		db:     db,
		config: appConfig,
	}
}
