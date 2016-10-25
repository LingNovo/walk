package core

import (
	"os"
	"testing"
)

// Walk 功能测试
func Test_Walk(t *testing.T) {
	e := Walk((os.Getenv("GOROOT")), "", "", (os.Getenv("GOROOT"))+"/out_test")
	if e != nil {
		t.Error(e)
	}
	//t.Log("over")
}

// Walk 基准测试
func Benchmark_Walk(b *testing.B) {
	e := Walk((os.Getenv("GOROOT")), "", "", (os.Getenv("GOROOT"))+"/out_test")
	if e != nil {
		b.Error(e)
	}
	//b.Log("over")
}
