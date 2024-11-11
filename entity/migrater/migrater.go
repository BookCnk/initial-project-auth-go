package migrater

import (
	"gorm.io/gorm"
	"initial-project-go/entity"
)

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entity.Project{},
		&entity.BugReport{},
		&entity.Permission{},
		&entity.Role{},
		&entity.User{},
		&entity.ApiKeys{},
		&entity.Encryptor{},
		&entity.UserToken{},
		&entity.UserAuth{},
	)
	if err != nil {
		return err
	}
	return nil
}
