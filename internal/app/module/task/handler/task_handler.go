package handler

import (
	"git.woa.com/ma/golang/utils/rsp"
	"github.com/gin-gonic/gin"
	"github.com/jakecoffman/cron"
	"github.com/jinzhu/gorm"
	"github.com/ruis09/go-task/internal/app/module/task/application"
	"strconv"
)

// TaskHandler
type TaskHandler struct {
	taskApp application.TaskApp
}

// NewTaskHandler
func NewTaskHandler(db *gorm.DB, cron *cron.Cron) *TaskHandler {
	return &TaskHandler{
		taskApp: application.NewTaskApp(db, cron),
	}
}

func (taskHandler *TaskHandler) Start(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		return
	}

	//调用app
	taskHandler.taskApp.Start(id)

	rsp.Success(c, "启动成功")
}

func (taskHandler *TaskHandler) Run(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		return
	}

	//调用app
	taskHandler.taskApp.Run(id)

	rsp.Success(c, "运行成功")
}

func (taskHandler *TaskHandler) Stop(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		return
	}

	taskHandler.taskApp.Stop(id)
	rsp.Success(c, "停止成功")
}

func (taskHandler *TaskHandler) Remove(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		return
	}

	taskHandler.taskApp.Remove(id)
	rsp.Success(c, "移除成功")
}
