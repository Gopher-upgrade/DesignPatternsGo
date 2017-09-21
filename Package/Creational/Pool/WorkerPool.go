package Pool

import (
	"fmt"
	"strconv"
	"time"
	"math/rand"
)

//声明成游戏
type Payload struct {
	name string
}

//打游戏
func (p *Payload) Play() {
	fmt.Printf("%s 打LOL游戏...当前任务完成\n", p.name)
}

//任务
type Job struct {
	Payload
}

//任务队列
var JobQueue chan Job

//  工人
type Worker struct {
	name string //工人的名字
	// WorkerPool chan JobQueue //对象池
	WorkerPool chan chan Job //对象池
	JobChannel chan Job      //通道里面拿
	quit       chan bool     //
}

// 新建一个工人
func NewWorker(workerPool chan chan Job, name string) Worker {

	fmt.Printf("创建了一个工人,它的名字是:%s \n", name);
	return Worker{
		name:       name,           //工人的名字
		WorkerPool: workerPool,     //工人在哪个对象池里工作,可以理解成部门
		JobChannel: make(chan Job), //工人的任务
		quit:       make(chan bool),
	}
}

// 工人开始工作
func (w *Worker) Start() {
	//开一个新的协程
	go func() {
		for {
			//注册到对象池中,
			w.WorkerPool <- w.JobChannel
			fmt.Printf("[%s]把自己注册到 对象池中 \n", w.name)
			select {
			//接收到了新的任务
			case job := <-w.JobChannel:
				fmt.Printf("[%s] 工人接收到了任务 当前任务的长度是[%d]\n", w.name, len(w.WorkerPool))
				job.Payload.Play()
				time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
				//接收到了任务
			case <-w.quit:
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

type Dispatcher struct {
	//WorkerPool chan JobQueue
	name       string        //调度的名字
	maxWorkers int           //获取 调试的大小
	WorkerPool chan chan Job //注册和工人一样的通道
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{
		WorkerPool: pool,       // 将工人放到一个池中,可以理解成一个部门中
		name:       "调度者",      //调度者的名字
		maxWorkers: maxWorkers, //这个调度者有好多个工人
	}
}

func (d *Dispatcher) Run() {
	// 开始运行
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.WorkerPool, fmt.Sprintf("work-%s", strconv.Itoa(i)))
		//开始工作
		worker.Start()
	}
	//监控
	go d.dispatch()

}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			fmt.Println("调度者,接收到一个工作任务")
			time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
			// 调度者接收到一个工作任务
			go func(job Job) {
				//从现有的对象池中拿出一个
				jobChannel := <-d.WorkerPool

				jobChannel <- job

			}(job)
		default:
			//fmt.Println("ok!!")
		}

	}
}
