package entity

import (
	"image"
)

// UploadFromType 图片上传来源
type UploadFromType int

const (
	UploadFromLocalFile UploadFromType = iota + 1 // 本地文件
	UploadFromWebUrl                              // 网络文件
	UploadFromAI                                  // AI生成
)

// StorageMethod 图片存储类型
type StorageMethod int

const (
	StorageToLocalFile StorageMethod = iota + 1 // 存储在本地分布式存储系统中
	StorageToCosFile                            // 存在在COS服务中
)

// UploadInfo 上传信息
type UploadInfo struct {
	FileName   string         // 本地文件名称
	WebUrl     string         // 网络图片URL
	UploadFrom UploadFromType // 图片上传来源
	Method     StorageMethod
	Style      *Style
	Tag        *Tag
}

// UploadFile 上传文件
type UploadFile struct {
	UploadInfo *UploadInfo // 图片上传信息

	SaveMethod      StorageMethod // 保存方式
	OriginImagePath string        // 保存路径
	WaterImagePath  string        // 水印图
}

// WaterImage 水印图
type WaterImage struct {
	Image    image.Image
	Text     string
	LeftTopX uint64
	LeftTopY uint64
	Opacity  float64
}

// ImageAttribute 图片属性
type ImageAttribute struct {
	Width uint32
	High  uint32

	WaterImage *WaterImage
}
