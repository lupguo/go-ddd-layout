package repository

import (
	"github.com/lupguo/go-ddd-layout/app/domain/entity"
)

type IReposStorage interface {
	// SaveImage 图片存储到指定服务（可能是本地文件存储服务、COS等）
	SaveImage(imgs []*entity.Image) error

	// FindImages 基于图片ID从存储中找到文件
	FindImages(ids []uint64) ([]*entity.Image, error)
}
