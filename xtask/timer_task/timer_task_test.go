package timer_task

import (
	"log"
	"os"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	log.Println("...开始执行任务...")
	timeout := 2 * time.Second
	r := New(timeout)
	r.Add(createTask(0), createTask(1), createTask(2))
	if err := r.Start(); err != nil {
		switch err {
		case ErrInterruput:
			log.Println(err)
			os.Exit(2)
		default:
			break
		}
	}
	log.Println("...任务执行结束...")
}

func Test2(t *testing.T) {
	log.Println("...开始执行任务...")
	timeout := 2 * time.Second
	r := New(timeout)
	r.Add(createTask(0), createTask(1), createTask(2))
	r.StartAsync(func(err error) {
		if err := r.Start(); err != nil {
			switch err {
			case ErrInterruput:
				log.Println(err)
				os.Exit(2)
			default:
				break
			}
		}
		log.Println("...任务执行结束...")
	})

	time.Sleep(time.Duration(10) * time.Second)

}
func createTask(param int) func() {
	return func() {
		log.Printf("正在执行任务%d", param)
		time.Sleep(time.Duration(param*100) * time.Second)
	}
}
