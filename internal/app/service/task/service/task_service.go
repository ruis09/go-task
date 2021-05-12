package service

import (
	"github.com/jakecoffman/cron"
	"github.com/ruis09/go-task/internal/app/service/task/repository/model"
)

type TaskService interface {
	InitTask()
	Add(task model.Task)                    //添加到任务列表中
	CreateJob(task model.Task) cron.FuncJob //执行一次任务
	Stop(task model.Task)                   //停止一次定时任务
	Remove(task model.Task)                 //停止并从任务列表移除
	GetTaskById(id int) (*model.Task, error)
}
