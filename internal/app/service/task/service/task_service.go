package service

import (
	"github.com/jakecoffman/cron"
	"github.com/ruis09/go-task/internal/app/service/task/repository/model"
)

type TaskService interface {
	InitTask()
	CreateJob(task model.Task) cron.FuncJob
	Stop(task model.Task)
	Remove(task model.Task)
	GetTaskById(id string) (*model.Task, error)
}
