package service

import (
	"github.com/lupguo/ddd-layout/app/domain/entity"
	"github.com/lupguo/ddd-layout/app/domain/repository"
)

type IServiceSearch interface {
	Search(keywords string) []*entity.UploadImage
}

type ImageSearchSrv struct{
	esRepos *repository.IReposSearch
	dbRepos *repository.IReposStorage
}

func (s *ImageSearchSrv) Search(keywords string) []*entity.UploadImage {
	return nil
}
