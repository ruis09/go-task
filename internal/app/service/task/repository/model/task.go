package model

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	Cron    string `gorm:"column:task_cron;"`   // cron表达式
	Name    string `gorm:"column:task_name"`    // 任务名称
	Command string `gorm:"column:task_command"` // 任务名称
	TimeOut string `gorm:"column:timeout"`      // 超时时间
	Status  string `gorm:"column:status"`       // 状态
}

func (m *Task) TableName() string {
	return "task"
}
