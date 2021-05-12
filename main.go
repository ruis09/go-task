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

	_, _ = task.Call("task1")
	_, _ = task.Call("task2")

	routes.Register(g, db, serviceCron)

	_ = g.Run("127.0.0.1:8888")

	////初始化定时任务
	//service.Initialize()
	//
	//r.GET("/stop", func(c *gin.Context) {
	//	name := c.Query("name")
	//	service.Stop(name)
	//	service.Remove(name)
	//
	//	c.JSON(200, gin.H{
	//		"result" : "stop success",
	//	})
	//})
	//r.GET("/remove", func(c *gin.Context) {
	//	name := c.Query("name")
	//	service.Remove(name)
	//
	//	c.JSON(200, gin.H{
	//		"result" : "stop success",
	//	})
	//})
	//r.GET("/start", func(c *gin.Context) {
	//	service.Run(service.Task{Cron:"* * * * * *", Name:"任务3", TimeOut:6})
	//
	//	c.JSON(200, gin.H{
	//		"result" : "start success",
	//	})
	//})
	//
	//_ = r.Run("127.0.0.1:8888")
}
