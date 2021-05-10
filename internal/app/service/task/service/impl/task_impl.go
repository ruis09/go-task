package impl

import (
	"github.com/jakecoffman/cron"
	"github.com/jinzhu/gorm"
	"github.com/ruis09/go-task/internal/app/service/task/repository"
	"github.com/ruis09/go-task/internal/app/service/task/repository/model"
	"github.com/ruis09/go-task/internal/app/service/task/repository/persistence"
	"golang.org/x/net/context"
	"sync"
	"time"
)

// TaskServiceImpl
type TaskServiceImpl struct {
	taskRepo repository.TaskRepo
	cron     *cron.Cron
}

// NewTaskServiceImpl
func NewTaskService(db *gorm.DB, cron *cron.Cron) *TaskServiceImpl {
	return &TaskServiceImpl{
		taskRepo: persistence.NewTaskRepo(db),
		cron:     cron,
	}
}

// 任务计数
type TaskCount struct {
	wg   sync.WaitGroup
	exit chan struct{}
}

func (tc *TaskCount) Add() {
	tc.wg.Add(1)
}

func (tc *TaskCount) Done() {
	tc.wg.Done()
}

func (tc *TaskCount) Exit() {
	tc.wg.Done()
	<-tc.exit
}

func (tc *TaskCount) Wait() {
	tc.Add()
	tc.wg.Wait()
	close(tc.exit)
}

// 并发队列
type ConcurrencyQueue struct {
	queue chan struct{}
}

func (cq *ConcurrencyQueue) Add() {
	cq.queue <- struct{}{}
}

func (cq *ConcurrencyQueue) Done() {
	<-cq.queue
}

var (
	taskCount TaskCount

	// 并发队列, 限制同时运行的任务数量
	concurrencyQueue ConcurrencyQueue

	taskMap sync.Map
)

// 初始化任务列表
func (t *TaskServiceImpl) InitTask() {

	concurrencyQueue = ConcurrencyQueue{queue: make(chan struct{}, 20)}

	taskCount = TaskCount{sync.WaitGroup{}, make(chan struct{})}
	go taskCount.Wait()

	lists, _ := t.taskRepo.GetTaskList()

	for _, task := range lists {
		taskFunc := t.CreateJob(task)
		t.cron.AddFunc(task.Cron, taskFunc, task.Name)
	}
}

type Result struct {
	string
}

func (t *TaskServiceImpl) CreateJob(task model.Task) cron.FuncJob {
	return func() {
		taskCount.Add()
		defer taskCount.Done()

		//TODO 开始执行之前记录日志

		concurrencyQueue.Add()
		defer concurrencyQueue.Done()

		//执行任务
		go func(t model.Task) {

			timeout := time.Duration(1) * time.Second
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()

			taskMap.Store(t.Name, cancel)
			defer taskMap.Delete(t.Name)

			println("正在执行：", t.Name)
			//exec(ctx, t)
			if t.Name == "任务2" {
				time.Sleep(time.Second * 5)
			}

			resultChan := make(chan Result)

			go func() {
				//todo 具体任务执行

				resultChan <- Result{"done"}
			}()

			select {
			case <-ctx.Done():
				println("任务停止：", t.Name)
			case <-resultChan:
				println("执行完毕：", t.Name)
			}

		}(task)

		//TODO 执行完记录日志

	}
}

func (t *TaskServiceImpl) Stop(task model.Task) {
	//先stop
	cancel, ok := taskMap.Load(task.Name)
	if !ok {
		return
	}
	cancel.(context.CancelFunc)()

	//最后remove
	t.Remove(task)
}

func (t *TaskServiceImpl) Remove(task model.Task) {
	t.cron.RemoveJob(task.Name)
}

func (t *TaskServiceImpl) GetTaskById(id string) (*model.Task, error) {
	return t.taskRepo.GetTaskById(id)
}
