package entity

import (
	"os"
)

// UploadFile 上传文件
type UploadFile struct {
	File *os.File
}

