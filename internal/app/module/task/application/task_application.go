package application

import (
	"github.com/jakecoffman/cron"
	"github.com/jinzhu/gorm"
	"github.com/ruis09/go-task/internal/app/service/task/service"
	"github.com/ruis09/go-task/internal/app/service/task/service/impl"
)

type TaskApp interface {
	Run(id string)
	Stop(id string)
	Remove(id string)
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

func (t *TaskAppImpl) Run(id string) {
	task, _ := t.taskSrv.GetTaskById(id)
	go t.taskSrv.CreateJob(*task)()
}

func (t *TaskAppImpl) Stop(id string) {
	task, _ := t.taskSrv.GetTaskById(id)
	t.taskSrv.Stop(*task)
}

func (t *TaskAppImpl) Remove(id string) {
	task, _ := t.taskSrv.GetTaskById(id)
	t.taskSrv.Remove(*task)
}
