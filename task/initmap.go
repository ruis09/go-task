package task

import "github.com/ruis09/go-task/task/job"

func InitTaskMap() {
	StubStorage = map[string]interface{}{
		"task1": job.RunTaskOne,
		"task2": job.RunTaskTwo,
	}
}
