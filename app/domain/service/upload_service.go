package service

import (
	"os"

	"github.com/lupguo/ddd-layout/app/domain/entity"
	"github.com/lupguo/ddd-layout/app/domain/repository"
)

type IServiceUpload interface {
	UploadImage(uploads []*os.File) ([]*entity.UploadImage, error)
	FindUploadImage(ids []uint64) ([]*entity.UploadImage, error)
}

type UploadImageSrv struct {
	dbRepos repository.IReposStorage
}

// UploadImage 图片上传服务
func (s *UploadImageSrv) UploadImage(uploads []*os.File) ([]*entity.UploadImage, error) {
	var entUploadImgs []*entity.UploadImage
	return s.dbRepos.SaveImage(entUploadImgs)
}

// FindUploadImage 查询上传图片
func (s *UploadImageSrv) FindUploadImage(ids []uint64) ([]*entity.UploadImage, error) {
	return s.dbRepos.FindImage(ids)
}

