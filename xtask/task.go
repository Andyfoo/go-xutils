package xtask

import (
	"os"
	"os/signal"
	"sync"
	"time"
)

type FinishType uint
type TimeOut time.Duration

const (
	//任务执行超时
	FinishType_TimeOut FinishType = iota
	//任务执行中断
	FinishType_Interruput
	//任务结束
	FinishType_Finish
)
const (
	TimeOut_Max   = time.Duration(100*365*24) * time.Hour
	TimeOut_Year  = time.Duration(365*24) * time.Hour
	TimeOut_Month = time.Duration(30*24) * time.Hour
	TimeOut_Day   = time.Duration(24) * time.Hour
	TimeOut_Hour  = time.Hour
)

type Runner struct {
	tasks     []func()         //要执行的任务
	complete  chan FinishType  //用于通知任务全部完成
	timeout   <-chan time.Time //这些任务在多久内完成 只能接收
	interrupt chan os.Signal   //可以控制强制终止的信号
}

func New(tm time.Duration) *Runner {
	return &Runner{
		complete:  make(chan FinishType), //同步通道，main routine等待，一直要任务完成或者被强制终止
		timeout:   time.After(tm),
		interrupt: make(chan os.Signal, 1), //至少接收到一个操作系统的中断信息
	}
}

//
func (r *Runner) Add(tasks ...func()) {
	r.tasks = append(r.tasks, tasks...)
}

//
func (r *Runner) run() FinishType {
	for _, task := range r.tasks {
		if r.isInterrupt() {
			return FinishType_Interruput
		}
		if r.finish {
			return FinishType_Finish
		}
		task()
	}
	return FinishType_Finish
}

//检查是否接收到了中断信号
func (r *Runner) isInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}

//开始执行所有任务，并且监视通道事件
func (r *Runner) Start() FinishType {
	//希望接收哪些系统信号
	signal.Notify(r.interrupt, os.Interrupt) //如果有系统中断的信号，发给r.interrupt
	var wg sync.WaitGroup
	go func() {
		r.complete <- r.run()
	}()

	select {
	case finishType := <-r.complete:
		r.finish = true
		return finishType
	case <-r.timeout:
		r.finish = true
		return FinishType_TimeOut
	}
}

//异步执行
func (r *Runner) StartAsync(callback func(finishType FinishType)) {
	go func() {
		finishType := r.Start()
		if callback != nil {
			callback(finishType)
		}
	}()
}
