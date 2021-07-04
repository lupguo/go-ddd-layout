package entity

import (
	"image"
)

// WaterImage 水印图
type WaterImage struct {
	Image    image.Image
	Text     string
	LeftTopX uint64
	LeftTopY uint64
	Opacity  float64
}
