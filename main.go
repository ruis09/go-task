package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jakecoffman/cron"
	"github.com/jinzhu/gorm"
	"github.com/ruis09/go-task/internal/app/service/task/service/impl"
	"github.com/ruis09/go-task/routes"
	"github.com/ruis09/go-task/task"
	"github.com/ruis09/go-task/tool"
)

var db *gorm.DB
var serviceCron *cron.Cron

type Exec struct {
}

func (e *Exec) run() {
	println("exec")
}

func main() {
	g := gin.New()

	db = tool.StartDB()

	serviceCron = cron.New()
	serviceCron.Start()

	//初始化任务方法
	task.InitTaskMap()

	//启动任务
	impl.NewTaskService(db, serviceCron).InitTask()

	routes.Register(g, db, serviceCron)

	_ = g.Run("127.0.0.1:8888")
}
