// Copyright 2019 Andyfoo
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package xthread

import (
	"sync"
)

type TaskFun func() interface{}
type TaskChanFun func(ch1 chan interface{})
type TaskFinishFun func(data interface{})
type AllFinishFun func()

type Thread struct {
	taskFuns     []TaskFun
	taskChanFuns []TaskChanFun
	waitGroup    sync.WaitGroup
}

func NewThread() Thread {
	return Thread{}
}

func (t *Thread) addTaskFuns(fun TaskFun) {
	t.taskFuns = append(t.taskFuns, fun)
}
func (t *Thread) addTaskChanFuns(fun TaskChanFun) {
	t.taskChanFuns = append(t.taskChanFuns, fun)
}
func (t *Thread) Start(taskCallback TaskFinishFun, allFinishCallback AllFinishFun) {
	ch1 := make(chan interface{})
	for _, fun1 := range t.taskFuns {
		//xlog.Info("run task:", k)
		t.waitGroup.Add(1)
		go func() {
			defer t.waitGroup.Done()
			ch1 <- fun1()
		}()
	}
	for _, fun1 := range t.taskChanFuns {
		//xlog.Info("run chan task:", k)
		t.waitGroup.Add(1)
		go func() {
			defer t.waitGroup.Done()
			fun1(ch1)
		}()
	}
	go func() {
		t.waitGroup.Wait()
		close(ch1)
	}()
	go func() {
		if taskCallback == nil {
			for _ = range ch1 {

			}
		} else {

			for data := range ch1 {
				taskCallback(data)
			}
		}

		if allFinishCallback != nil {
			allFinishCallback()
		}
	}()

}
func (t *Thread) StartSync(taskCallback TaskFinishFun) {
	ch1 := make(chan interface{})
	for _, fun1 := range t.taskFuns {
		//xlog.Info("run task:", k)
		t.waitGroup.Add(1)
		go func() {
			defer t.waitGroup.Done()
			ch1 <- fun1()
		}()
	}
	for _, fun1 := range t.taskChanFuns {
		//xlog.Info("run chan task:", k)
		t.waitGroup.Add(1)
		go func() {
			defer t.waitGroup.Done()
			fun1(ch1)
		}()
	}
	go func() {
		t.waitGroup.Wait()
		close(ch1)
	}()
	if taskCallback == nil {
		for _ = range ch1 {

		}
	} else {

		for data := range ch1 {
			taskCallback(data)
		}
	}
}
