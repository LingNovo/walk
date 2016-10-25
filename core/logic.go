package core

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

/*
* 描述: 搜索指定目录下的文件，并记录文件信息
* 参数列表
* @path: 表示根目录字符
* @matchDir： 表示目录过滤字符，包含该字符的会被忽略
* @matchFile: 表示文件过滤字符，包含该字符的会被忽略
* @outFile: 表示记录信息的文件路径
* 返回值列表:
* 返回错误信息
 */
func Walk(path, matchDir, matchFile, outFile string) error {
	dirReg, e := strToReg(matchDir)
	if e != nil {
		return e
	}
	fileReg, e := strToReg(matchFile)
	if e != nil {
		return e
	}
	var array []string
	e = filepath.Walk(path, func(root string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if dirReg != nil && dirReg.MatchString(root) {
			return nil
		}
		if fileReg != nil && fileReg.MatchString(info.Name()) {
			return nil
		}
		array = append(array, root)
		return nil
	})
	if e != nil {
		return e
	}
	r := newRecorder(outFile)
	pool := new(goPool)
	if len(array) > 1024 { //linux 打开文件数上限1024,超过会报错。
		pool.Init(1000, len(array))
	} else {
		pool.Init(len(array)/2, len(array))
	}
	for i := range array {
		path := array[i]
		pool.AddTask(func() error {
			f, e := newFileInfoEx(path)
			if e != nil {
				return e
			}
			_, e = r.Write([]byte(f.format()))
			return e
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
	r.Close()
	pool.Stop()
	return nil
}

// 字符转正则，拼接过滤符
func strToReg(str string) (*regexp.Regexp, error) {
	if len(strings.TrimSpace(str)) == 0 {
		return nil, nil
	}
	if strings.Index(str, "*") == 0 {
		str = "." + str
	} else {
		str = "^" + str
	}
	str += "$"
	reg, err := regexp.Compile(str)
	if err != nil {
		return nil, err
	}
	return reg, nil
}
