package repository

import (
	"github.com/lupguo/go-ddd-layout/app/domain/entity"
)

type IReposSearch interface {
	Search(req *entity.SearchKeywords) []*entity.SearchImage
	// SearchTags(req *entity.SearchTagReq) []*SearchImage
}
