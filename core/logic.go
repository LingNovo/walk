package core

import (
	"os"
	"path/filepath"
	"strings"
)

/*
* 描述: 获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。
* 参数列表:
* @root: 表示根目录路径字符串
* @ignore: 表示可被忽略的目录的关键词字符串,eg:"fx",表示包含"fx"字符的目录均被忽略
* @suffix: 表示匹配文件后缀字符,eg:".pdf"
* @outFilePath: 表示输出结果文件路径
* 返回值列表:
* 返回error错误信息
 */
func Walk(root, ignore, suffix, outFilePath string) error {
	var (
		err     error
		fInfoEx *FileInfoEx
	)
	err = filepath.Walk(root, func(fName string, fInfo os.FileInfo, e error) error {
		if e != nil {
			return e
		}
		if fInfo.IsDir() {
			return nil
		}
		if len(ignore) > 0 && strings.Contains(strings.ToUpper(filepath.Dir(fName)), strings.ToUpper(ignore)) {
			return nil
		}
		if len(suffix) > 0 && !strings.HasSuffix(strings.ToUpper(fName), strings.ToUpper(suffix)) {
			return nil
		}
		if fInfoEx, e = NewFileInfoEx(fName); e != nil {
			return e
		}
		fInfoEx.writeToFile(outFilePath)
		return nil
	})
	return err
}
