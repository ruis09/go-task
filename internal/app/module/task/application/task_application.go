package application

import (
	"github.com/jakecoffman/cron"
	"github.com/jinzhu/gorm"
	"github.com/ruis09/go-task/internal/app/service/task/service"
	"github.com/ruis09/go-task/internal/app/service/task/service/impl"
)

type TaskApp interface {
	Start(id int)
	Run(id int)
	Stop(id int)
	Remove(id int)
}

// TaskAppImpl
type TaskAppImpl struct {
	taskSrv service.TaskService
}

// NewTaskApp
func NewTaskApp(db *gorm.DB, cron *cron.Cron) *TaskAppImpl {
	return &TaskAppImpl{
		taskSrv: impl.NewTaskService(db, cron),
	}
}

func (t *TaskAppImpl) Start(id int) {
	task, _ := t.taskSrv.GetTaskById(id)
	go t.taskSrv.Add(*task)
}

func (t *TaskAppImpl) Run(id int) {
	task, _ := t.taskSrv.GetTaskById(id)
	go t.taskSrv.CreateJob(*task)()
}

func (t *TaskAppImpl) Stop(id int) {
	task, _ := t.taskSrv.GetTaskById(id)
	t.taskSrv.Stop(*task)
	t.taskSrv.Remove(*task)
}

func (t *TaskAppImpl) Remove(id int) {
	task, _ := t.taskSrv.GetTaskById(id)
	t.taskSrv.Remove(*task)
}
