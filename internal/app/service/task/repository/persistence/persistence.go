package persistence

import "github.com/jinzhu/gorm"

type TaskRepoImpl struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) *TaskRepoImpl {
	return &TaskRepoImpl{
		db: db,
	}
}
