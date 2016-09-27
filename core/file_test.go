package core

import (
	"testing"
)

// 方法 NewFileInfoEx 功能测试
func Test_NewFileInfoEx(t *testing.T) {
	var fPath string = "d:\\develop\\res\\分布式算法导论.pdf"
	if _, e := NewFileInfoEx(fPath); e != nil {
		t.Error(e)
	}
}

// 方法 NewFileInfoEx 基准测试
func Benchmark_NewFileInfoEx(b *testing.B) {
	var fPath string = "d:\\develop\\res\\分布式算法导论.pdf"
	if _, e := NewFileInfoEx(fPath); e != nil {
		b.Error(e)
	}
}
