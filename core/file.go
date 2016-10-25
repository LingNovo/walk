package core

import (
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

// 文件信息结构
type fileInfoEx struct {
	// 文件名称
	Name string
	// 文件哈希值
	Hash string
	// 文件大小
	Size int64
}

/*
* 描述: 格式化文件信息fileInfoEx实例为字符串
* 返回值列表:
* 返回一个格式化后的fileInfoEx字符串
 */
func (this *fileInfoEx) format() string {
	return fmt.Sprintf("名称:%s,哈希值:%s,大小:%d\r\n", this.Name, this.Hash, this.Size)
}

/*
* 描述: 实例化一个fileInfoEx文件信息
* 参数列表:
* @fullPath: 表示一个文件路径字符串
* 返回值列表:
* @info: 表示一个fileInfoEx文件信息指针对象
* @e: 表示错误信息
 */
func newFileInfoEx(fullPath string) (*fileInfoEx, error) {
	f, e := os.OpenFile(fullPath, os.O_RDONLY, os.ModePerm|os.ModeTemporary)
	defer f.Close()
	if e != nil {
		return nil, e
	}

	r := fileInfoEx{}
	info, e := f.Stat()
	if e != nil {
		return nil, e
	}
	r.Name = info.Name()
	r.Size = info.Size()
	h := sha1.New()
	_, e = io.Copy(h, f)
	if e != nil {
		return nil, e
	}
	r.Hash = fmt.Sprintf("%x", h.Sum(nil))
	return &r, nil
}
