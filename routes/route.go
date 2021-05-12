package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jakecoffman/cron"
	"github.com/jinzhu/gorm"
	"github.com/ruis09/go-task/internal/app/module/task/handler"
)

func Register(g *gin.Engine, db *gorm.DB, cron *cron.Cron) *gin.Engine {
	// V1 API
	v1 := g.Group("/api/v1/task")
	{
		task := initTaskHandler(db, cron)
		{
			v1.POST("/start/:task_id", task.Start)   //启动任务
			v1.POST("/run/:task_id", task.Run)       //运行一次
			v1.POST("/stop/:task_id", task.Stop)     //停止一次任务
			v1.POST("/remove/:task_id", task.Remove) //移除任务
		}
	}

	return g
}

// 初始化任务Handler
func initTaskHandler(db *gorm.DB, cron *cron.Cron) *handler.TaskHandler {
	return handler.NewTaskHandler(db, cron)
}
