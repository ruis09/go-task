package task

import "github.com/ruis09/go-task/task/job"

func InitTaskMap() {
	StubStorage = map[string]interface{}{
		"任务1": job.RunTaskOne,
		"任务2": job.RunTaskTwo,
	}
}
