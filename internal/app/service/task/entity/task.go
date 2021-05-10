package entity

type Task struct {
	ID      int
	Cron    string `json:"task_cron"`    // cron表达式
	Name    string `json:"task_name"`    // 任务名称
	Command string `json:"task_command"` //
	TimeOut string `json:"timeout"`      //超时时间
	Status  string `json:"status"`       //状态
}
