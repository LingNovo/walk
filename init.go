package main

import (
	"flag"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	// 根目录，默认值为当前应用程序工作目录
	Root *string
	// 可被忽略的目录的关键词，默认值为空，即不过滤
	Ignore *string
	// 匹配文件后缀字符，默认值为空，即不过滤
	Suffix *string
	// 输出结果文件路径，默认值为当前应用程序工作目录下
	OutFilePath string
)

func init() {
	var (
		e          error
		currentDir string
		t          time.Time = time.Now()
		f          *os.File
	)
	// 获取当前目录
	if currentDir, e = filepath.Abs(filepath.Dir(os.Args[0])); e != nil {
		panic(e)
	}
	currentDir = strings.Replace(currentDir, "\\", "/", -1)
	// 接收控制台参数及初始化
	Root = flag.String("root", currentDir, "Use -root <root dir source>")
	Ignore = flag.String("ignore", "", "Use -ignore <ignore dir filter>")
	Suffix = flag.String("suffix", "", "Use -suffix <file suffix filter>")
	flag.Parse()
	// 创建记录结果的文件
	OutFilePath = string(*Root) + "/record_" + strings.Replace(t.String()[:19], ":", "_", 3) + ".txt"
	if f, e = os.OpenFile(OutFilePath, os.O_CREATE, 0666); e != nil {
		panic(e)
	}
	defer f.Close()

}
