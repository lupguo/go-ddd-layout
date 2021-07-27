package entity

import (
	"image"
)

// UploadImage 图片文件
type UploadImage struct {
	ID    uint64
	Image image.Image
	Url   string
	Tags  []*Tags
}

// ImageAttribute 图片属性
type ImageAttribute struct {
	Width uint32
	High  uint32

	WaterImage *WaterImage
}
