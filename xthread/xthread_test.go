package xthread

import (
	"fmt"
	"testing"
	"time"

	"github.com/Andyfoo/go-xutils/xvar"
)

func _Test1(t *testing.T) {
	fmt.Println("task")
	//var mylist = []string{}
	t1 := NewThread()
	t1.AddTaskFuns(func() interface{} {
		return "data1"
	})
	t1.AddTaskFuns(func() interface{} {
		return "data2"
	})
	t1.AddTaskChanFuns(func(ch1 chan interface{}) {
		for i := 0; i < 2; i++ {
			ch1 <- xvar.Int64ToStr(time.Now().UnixNano())
		}
	})
	t1.StartSync(func(data interface{}) {
		fmt.Println(data)
	})
	// t1.Start(func(data interface{}) {
	// 	fmt.Println(data)
	// }, func() {
	// 	fmt.Println("finish 1")
	// })
	fmt.Println("finish")
	time.Sleep(100 * time.Microsecond)
}
func Test2(t *testing.T) {
	fmt.Println("task")
	//var mylist = []string{}
	t1 := NewThread()
	for i := 0; i < 10; i++ {

		func(i2 int) {
			t1.AddTaskChanFuns(func(ch1 chan interface{}) {
				fmt.Println(i2)
				for i := 0; i < 2; i++ {
					ch1 <- xvar.Int64ToStr(time.Now().UnixNano())
				}
			})
		}(i)
	}

	t1.StartSync(func(data interface{}) {
		//fmt.Println(data)
	})

	fmt.Println("finish")
	time.Sleep(100 * time.Microsecond)
}
