package service

import (
	"github.com/lupguo/go-ddd-layout/app/domain/entity"
	"github.com/lupguo/go-ddd-layout/app/domain/repository"
)

// IImageSearchService 图片搜索服务
type IImageSearchService interface {

	// Search 通过图片名称、类目、标签、id等关键字搜索图片
	Search(keywords string) []*entity.Image
}

// ImageSearchSrv 搜索服务具体实现
type ImageSearchSrv struct {
	esRepos *repository.IReposSearch
	dbRepos *repository.IReposStorage
}

// NewImageSearchSrv 创建搜索服务
func NewImageSearchSrv(esRepos *repository.IReposSearch, dbRepos *repository.IReposStorage) *ImageSearchSrv {
	return &ImageSearchSrv{
		esRepos: esRepos,
		dbRepos: dbRepos,
	}
}

// Search 搜索关键字
func (s *ImageSearchSrv) Search(keywords string) []*entity.Image {
	return nil
}
