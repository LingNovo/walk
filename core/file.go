package core

import (
	"crypto/sha1"
	"fmt"
	"hash"
	"io"
	"os"
)

type (
	// 文件信息结构
	FileInfoEx struct {
		// 文件名称
		Name string
		// 文件大小
		Size int64
		// 文件哈希值
		Hash string
	}
)

/*
* 描述: 格式化文件信息FileInfoEx实例为字符串
* 返回值列表:
* 返回一个格式化后的FileInfoEx字符串
 */
func (this *FileInfoEx) format() string {
	return fmt.Sprintf("名称:%s,哈希值:%s,大小:%d\r\n", this.Name, this.Hash, this.Size)
}

/*
* 描述:获取文件哈希值
* 参数列表:
* @r: 表示一个io.reader 文件读取器
* 返回值列表:
* @code: 表示一个哈希值字符串
* @e: 表示错误信息
 */
func (this *FileInfoEx) getHash(r io.Reader) (code string, e error) {
	var h hash.Hash = sha1.New()
	if _, e = io.Copy(h, r); e != nil {
		return "", e
	}
	code = fmt.Sprintf("%x", h.Sum(nil))
	return code, nil
}

/*
* 描述: 将格式化后的FileInfoEx 信息写入指定文件
* 参数列表:
* @fPath: 表示一个文件路径字符串
* 返回值列表:
* 返回error错误信息
 */
func (this *FileInfoEx) writeToFile(fPath string) error {
	var (
		e error
		f *os.File
	)
	//打开一个文件，状态为只写，追加模式
	if f, e = os.OpenFile(fPath, os.O_APPEND|os.O_WRONLY, 0644); e != nil {
		return e
	}
	defer f.Close()
	f.WriteString(this.format())
	return nil
}

/*
* 描述: 实例化一个FileInfoEx文件信息
* 参数列表:
* @fPath: 表示一个文件路径字符串
* 返回值列表:
* @info: 表示一个FileInfoEx文件信息指针对象
* @e: 表示错误信息
 */
func NewFileInfoEx(fPath string) (info *FileInfoEx, e error) {
	var (
		f     *os.File
		fInfo os.FileInfo
	)
	if f, e = os.Open(fPath); e != nil {
		return nil, e
	}
	defer f.Close()
	info = &FileInfoEx{}
	if fInfo, e = f.Stat(); e != nil {
		return nil, e
	}
	info.Name = fInfo.Name()
	info.Size = fInfo.Size()
	if info.Hash, e = info.getHash(f); e != nil {
		return nil, e
	}
	return info, nil
}
