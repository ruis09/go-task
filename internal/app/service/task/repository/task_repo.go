package repository

import "github.com/ruis09/go-task/internal/app/service/task/repository/model"

type TaskRepo interface {
	GetTaskList() ([]model.Task, error)
	GetTaskById(id int) (*model.Task, error)
}
