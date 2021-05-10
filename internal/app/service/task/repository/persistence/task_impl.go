package persistence

import (
	"github.com/ruis09/go-task/internal/app/service/task/repository/model"
)

//操作数据库

func (t *TaskRepoImpl) GetTaskList() ([]model.Task, error) {
	var lists []model.Task
	var err error

	err = t.db.Model(model.Task{}).Find(&lists).Error

	return lists, err
}

func (t *TaskRepoImpl) GetTaskById(id string) (*model.Task, error) {
	var task model.Task
	var err error

	if err = t.db.First(&task).Error; err != nil {
		return nil, err
	}

	return &task, err
}
