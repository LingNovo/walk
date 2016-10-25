package core

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

// newFileInfoEx 功能测试
func Test_newFileInfoEx(t *testing.T) {
	array, e := ioutil.ReadDir(os.Getenv("GOROOT"))
	if e != nil {
		t.Error(e)
	}
	for _, v := range array {
		if !v.IsDir() && strings.HasSuffix(v.Name(), ".go") {
			path := os.TempDir() + "/" + v.Name()
			_, e := newFileInfoEx(path)
			if e != nil {
				t.Error(e)
			}
			//t.Log(f)
			break
		}
	}
}

// newFileInfoEx 基准测试
func Benchmark_newFileInfoEx(b *testing.B) {
	array, e := ioutil.ReadDir(os.Getenv("GOROOT"))
	if e != nil {
		b.Error(e)
	}
	for _, v := range array {
		if !v.IsDir() && strings.HasSuffix(v.Name(), ".go") {
			path := os.TempDir() + "/" + v.Name()
			_, e := newFileInfoEx(path)
			if e != nil {
				b.Error(e)
			}
			//b.Log(f)
			break
		}
	}
}
