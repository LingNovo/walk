package core

import (
	"fmt"
)

// 线程池
type goPool struct {
	Queue          chan func() error
	Number         int
	Total          int
	result         chan error
	finishCallback func()
}

// 初始化
func (this *goPool) Init(number, total int) {
	this.Queue = make(chan func() error, total)
	this.Number = number
	this.Total = total
	this.result = make(chan error, total)
}

// 开门迎客
func (this *goPool) Start() {
	for i := 0; i < this.Number; i++ {
		go func() {
			for {
				task, ok := <-this.Queue
				if !ok {
					break
				}
				err := task()
				this.result <- err
			}
		}()
	}
	for j := 0; j < this.Total; j++ {
		res, ok := <-this.result
		if !ok {
			break
		}
		if res != nil {
			fmt.Println(res)
		}
	}
	if this.finishCallback != nil {
		this.finishCallback()
	}
}

// 关门送客
func (this *goPool) Stop() {
	close(this.Queue)
	close(this.result)
}

// 添加任务
func (this *goPool) AddTask(task func() error) {
	this.Queue <- task
}

// 设置回调函数
func (this *goPool) SetFinshCallback(callback func()) {
	this.finishCallback = callback
}
