package core

import (
	"testing"
)

// 方法 Walk 功能测试
func Test_Walk(t *testing.T) {
	var (
		e           error
		root        string = "d:\\develop"
		ignore      string = "res"
		suffix      string = ".go"
		outFilePath string = "d:\\develop\\base\\src\\walk"
	)
	if e = Walk(root, ignore, suffix, outFilePath); e != nil {
		t.Error(e)
	}
}

// 方法 Walk 基准测试
func Benchmark_Walk(b *testing.B) {
	var (
		e           error
		root        string = "d:\\develop"
		ignore      string = "res"
		suffix      string = ".go"
		outFilePath string = "d:\\develop\\base\\src\\walk"
	)
	if e = Walk(root, ignore, suffix, outFilePath); e != nil {
		b.Error(e)
	}
}
