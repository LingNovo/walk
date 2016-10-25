package core

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

// 记录器功能测试
func Test_Write(t *testing.T) {
	array, e := ioutil.ReadDir((os.Getenv("GOROOT")))
	if e != nil {
		t.Error(e)
	}
	for _, v := range array {
		if !v.IsDir() && strings.HasSuffix(v.Name(), ".go") {
			path := os.TempDir() + "/" + v.Name()
			f, e := newFileInfoEx(path)
			if e != nil {
				t.Error(e)
			}
			out := os.TempDir() + "/out_test"
			r := newRecorder(out)
			_, e = r.Write([]byte(f.format()))
			if e != nil {
				t.Error(e)
			}
			break
		}
	}
}

// 记录器基准测试
func Benchmark_Write(b *testing.B) {
	array, e := ioutil.ReadDir((os.Getenv("GOROOT")))
	if e != nil {
		b.Error(e)
	}
	for _, v := range array {
		if !v.IsDir() && strings.HasSuffix(v.Name(), ".go") {
			path := os.TempDir() + "/" + v.Name()
			f, e := newFileInfoEx(path)
			if e != nil {
				b.Error(e)
			}
			out := os.TempDir() + "/out_test"
			r := newRecorder(out)
			_, e = r.Write([]byte(f.format()))
			if e != nil {
				b.Error(e)
			}
			break
		}
	}
}
