package core

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

//　线程池功能测试
func Test_Pool(t *testing.T) {
	array, e := ioutil.ReadDir((os.Getenv("GOROOT")))
	if e != nil {
		t.Error(e)
	}
	pool := new(goPool)
	if len(array) > 1024 { //linux 打开文件数上限1024,超过会报错。
		pool.Init(1000, len(array))
	} else {
		pool.Init(len(array)/2, len(array))
	}
	for i := range array {
		info := array[i]
		pool.AddTask(func() error {
			_, e := os.Stat(os.Getenv("GOROOT") + "/" + info.Name())
			if e != nil {
				t.Log(e)
				return e
			}
			return nil
		})
	}
	isFinsh := false
	pool.SetFinshCallback(func() {
		func(isFinsh *bool) {
			*isFinsh = true
		}(&isFinsh)
	})
	pool.Start()
	for !isFinsh {
		time.Sleep(time.Millisecond * 100)
	}
	pool.Stop()
	//t.Log("over")
}

// 线程池基准测试
func Benchmark_Pool(b *testing.B) {
	array, e := ioutil.ReadDir((os.Getenv("GOROOT")))
	if e != nil {
		b.Error(e)
	}
	pool := new(goPool)
	if len(array) > 1024 { //linux 打开文件数上限1024,超过会报错。
		pool.Init(1024, len(array))
	} else {
		pool.Init(len(array)/2, len(array))
	}
	for i := range array {
		info := array[i]
		pool.AddTask(func() error {
			_, e := os.Stat(os.Getenv("GOROOT") + "/" + info.Name())
			if e != nil {
				b.Log(e)
				return e
			}
			return nil
		})
	}
	isFinsh := false
	pool.SetFinshCallback(func() {
		func(isFinsh *bool) {
			*isFinsh = true
		}(&isFinsh)
	})
	pool.Start()
	for !isFinsh {
		time.Sleep(time.Millisecond * 100)
	}
	pool.Stop()
	//b.Log("over")
}
