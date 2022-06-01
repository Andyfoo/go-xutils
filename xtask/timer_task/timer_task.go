package timer_task

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

//任务执行超时
var ErrTimeOut = errors.New("run time out")

//任务执行中断
var ErrInterruput = errors.New("run interruput")

type TimerRunner struct {
	tasks     []func()         //要执行的任务
	complete  chan error       //用于通知任务全部完成
	timeout   <-chan time.Time //这些任务在多久内完成 只能接收
	interrupt chan os.Signal   //可以控制强制终止的信号
}

func New(tm time.Duration) *TimerRunner {
	return &TimerRunner{
		complete:  make(chan error), //同步通道，main routine等待，一致要任务完成或者被强制终止
		timeout:   time.After(tm),
		interrupt: make(chan os.Signal, 1), //至少接收到一个操作系统的中断信息
	}
}

//
func (r *TimerRunner) Add(tasks ...func()) {
	r.tasks = append(r.tasks, tasks...)
}

//
func (r *TimerRunner) run() error {
	for _, task := range r.tasks {
		if r.isInterrupt() {
			return ErrInterruput
		}
		task()
	}
	return nil
}

//检查是否接收到了中断信号
func (r *TimerRunner) isInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}

//开始执行所有任务，并且监视通道事件
func (r *TimerRunner) Start() error {
	//希望接收哪些系统信号
	signal.Notify(r.interrupt, os.Interrupt) //如果有系统中断的信号，发给r.interrupt

	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeOut
	}
}

//异步执行
func (r *TimerRunner) StartAsync(callback func(err error)) {
	go func() {
		err := r.Start()
		if callback != nil {
			callback(err)
		}
	}()
}
