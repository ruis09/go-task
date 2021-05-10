package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jakecoffman/cron"
	"github.com/jinzhu/gorm"
	"github.com/ruis09/go-task/internal/app/module/task/application"
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

	id := c.Param("id")

	//调用app
	taskHandler.taskApp.Run(id)
}

func (taskHandler *TaskHandler) Stop(c *gin.Context) {
	id := c.Param("id")

	taskHandler.taskApp.Stop(id)
}

func (taskHandler *TaskHandler) Remove(c *gin.Context) {
	id := c.Param("id")

	taskHandler.taskApp.Remove(id)
}
