package entity

type (
	Project struct {
		ProjectID   uint        `json:"id" gorm:"primaryKey; autoIncrement"`
		Name        string      `json:"name" gorm:"size:100;not null"`
		Description string      `json:"description"`
		BugReports  []BugReport `json:"bug_reports" gorm:"foreignKey:ProjectID"`
	}

	CreateProjectReq struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description" validate:"required"`
	}
)
