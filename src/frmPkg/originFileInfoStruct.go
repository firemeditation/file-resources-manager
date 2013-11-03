package frmPkg

import (
	"os"
	"time"
)

// OriginFileInfoStruct 源文件信息结构体，用于客户端获取本地需上传的文件的相关信息，并由服务端获得并解析。类似os.FileInfo的结构。
type OriginFileInfoStruct struct {
	RelativeDir string  //相对路径
	FileName string  //文件名(只是文件名)
	Size int64  //文件大小
	Mode os.FileMode  //文件权限
	ModeTime time.Time  //文件修改日期
}


type OriginFileInfoFullStruct struct {
	FullDir string  //完整客户端本地路径
	OriginFileInfoStruct
}
