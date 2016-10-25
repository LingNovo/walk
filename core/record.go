package core

import (
	"io"
	"os"
	"path/filepath"
)

// 记录器
type recorder struct {
	// 保存记录内容的文件名称
	fileName string
	// 带关闭的数据流写入器
	innerWriter io.WriteCloser
}

/*
* 描述: 实例化一个recorder记录器
* 参数列表:
* @fileName: 表示一个文件路径字符串
* 返回值列表:
* 返回表示一个recorder记录器指针对象
 */
func newRecorder(fileName string) *recorder {
	r := recorder{}
	r.fileName = fileName
	return &r
}

/*
* 描述: 创建记录文件，初始化记录器
* 返回值列表:
* @e: 表示错误信息
 */
func (this *recorder) createFile() (e error) {
	folder, _ := filepath.Split(this.fileName)
	if 0 != len(folder) {
		e = os.MkdirAll(folder, 0767)
		if e != nil {
			return e
		}
	}
	this.innerWriter, e = os.OpenFile(this.fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if e != nil {
		return e
	}
	return nil
}

/*
* 描述: 写入数据
* 参数列表:
* @data:　表示字节数组
* 返回值列表:
* @n: 写入的字节数
* @e: 表示错误信息
 */
func (this *recorder) Write(data []byte) (int, error) {
	if this.innerWriter == nil {
		if e := this.createFile(); e != nil {
			return 0, e
		}
	}
	return this.innerWriter.Write(data)
}

/*
* 描述: 关闭记录器，释放资源
* 返回值列表:
* 返回错误信息
 */
func (this *recorder) Close() (e error) {
	return this.innerWriter.Close()
}
